#!/bin/sh
until nc -z db 5432; do
  echo "Aguardando PostgreSQL..."
  sleep 1
done

echo "PostgreSQL está disponível. Iniciando aplicação..."
exec ./main
