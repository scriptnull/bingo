'use strict'

var express = require('express')
var bodyParser = require('body-parser')
var session = require('express-session')
var app = express()

var _ = require('underscore')

var passport = require('passport')
var FacebookStrategy = require('passport-facebook').Strategy

// load bingo config
try {
  var appConfig = require('../app.config.json')

  if (_.isEmpty(appConfig)) {
    throw new Error('Expected appConfig to be not empty')
  }

  if (_.isEmpty(appConfig['game-web'])) {
    throw new Error('Expected appConfig[\'game-web\'] to be not empty')
  }

  global.config = appConfig['game-web']
} catch (err) {
  console.error('Loading application configuration failed :: ' + err)
  process.exit(1)
}

// express server config
app.set('view engine', 'ejs')
app.use(express.static('public'))
app.use(session({
  secret: 'cat'
}))
app.use(bodyParser())
app.use(passport.initialize())
app.use(passport.session())

var tempUsers = {}

passport.serializeUser(
  function (user, done) {
    console.log('mylog serializeUser user ' + JSON.stringify(user))
    done(null, user.id)
  }
)

passport.deserializeUser(
  function (id, done) {
    console.log('mylog deserializeUser id ' + id)
    done(null, tempUsers[id])
  }
)

passport.use(new FacebookStrategy(
  {
    clientID: global.config.fb_app_id,
    clientSecret: global.config.fb_app_secret,
    callbackURL: global.config.fb_auth_callback
  },
  function (accessToken, refreshToken, profile, done) {
    tempUsers[profile.id] = profile
    done(null, profile)
  }
))

// express passport routes
app.get('/auth/facebook', passport.authenticate('facebook'))

app.get('/auth/facebook/callback',
  passport.authenticate('facebook',
    { successRedirect: '/', failureRedirect: '/login' }
  )
)

// express server custom routes
app.get('/',
  function (req, res) {
    if (_.isEmpty(req.user)) {
      return res.render('index')
    }

    return res.render('app')
  }
)

// start web server
app.listen(global.config.port, function () {
  console.log('game-web server started :: port ' + global.config.port)
})
