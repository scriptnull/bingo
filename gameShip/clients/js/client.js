var PROTO_PATH = __dirname + '/../../gameShipRpc/gameShipRpc.proto';

var fs = require('fs');
var grpc = require('grpc');
var proto = grpc.load(PROTO_PATH);

console.log(proto);

var gameShipRpc = proto.gameShipRpc;

console.log(gameShipRpc)

var client = new gameShipRpc.gameShipRpc('localhost:50051',
                                       grpc.credentials.createInsecure());

var gameId;

client.newGame({ creatorId: 'player-1'}, function (err, res) {
  console.log('*********************')
  console.log('New game created');
  console.log(err);
  console.log(res);
  gameId = res.gameId;
  client.playerJoin({ playerId: 'player-2', gameId: gameId}, function (err, res) {
    console.log('*********************')
    console.log('Player 2 joined')
    console.log(err);
    console.log(res);
    client.gameStart({ gameId: gameId}, function (err, res) {
      console.log('*********************')
      console.log('gameStarted')
      console.log(err);
      console.log(res);
      client.playerStrikeBox({ gameId: gameId, playerId: 'player-1', row: 1, column: 1}, function (err, res) {
        console.log('*********************')
        console.log('player-1 striked r1 c1')
        console.log(err);
        console.log(res);
        client.playerStrikeBox({ gameId: gameId, playerId: 'player-1', row: 1, column: 1}, function (err, res) {
          console.log('*********************')
          console.log('player-1 trying to strike again r1 c1 => should fail')
          console.log(err);
          console.log(res);
          client.playerStrikeBox({ gameId: gameId, playerId: 'player-2', row: 2, column: 2}, function (err, res) {
            console.log('*********************')
            console.log('player-2 strikes r2 c2 ')
            console.log(err);
            console.log(res);
            client.playerBingo({ gameId: gameId, playerId: 'player-2'}, function (err, res) {
              console.log('*********************')
              console.log('player-2 bingoes')
              console.log(err);
              console.log(res);
            })
          })
        })
      })
    })
  });
});
