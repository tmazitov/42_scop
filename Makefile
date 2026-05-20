NAME = scop

all: $(NAME)

$(NAME):
	go build -o $(NAME) ./cmd/

clean:
	rm -f $(NAME)

fclean: clean

re: fclean all

.PHONY: all clean fclean re
