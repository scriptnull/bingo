docker run --name bingo-postgres \
  -e POSTGRES_DB=bingodb \
  -e POSTGRES_USER=bingouser \
  -e POSTGRES_PASSWORD=bingopassword \
  -d postgres
