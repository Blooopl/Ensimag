SRCS = $(wildcard *.c)
OBJS := $(SRCS:%.c=%.o)
BINS := $(SRCS:%.c=%)

CC = clang
CFLAGS = -m64 -std=gnu99 -Wall -Wextra -Werror -g
# Pour que ça compile avec GCC qui a besoin que -lm soit après les sources qui s'en servent
LDLIBS = -lm

.PHONY: all
all: $(BINS)

.PHONY: clean
clean:
	$(RM) -r $(BINS) $(OBJS) $(foreach f, $(BINS), $(addsuffix .dSYM, $f))
