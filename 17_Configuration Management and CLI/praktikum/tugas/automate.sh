NAME=$1

if [ -z "$NAME" ]
then
	NAME=world
fi

echo "$NAME $(date)"
