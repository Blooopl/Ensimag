from itertools import repeat,chain


def repeter(iterable,k):
    for item in iterable:
        for indice in range(k):
            yield item

def repeter2(iterable,k):
    return chain(*(repeat(i,k) for i in iterable))

def paires(iterable):

    iterateur = iter(iterable)
    try:
        temp = next(iterateur)
    except StopIteration:
        return



    while True:

        try:
            temp2 = next(iterateur)
            yield (temp,temp2)
            temp = temp2


        except StopIteration:
            break



class IterateurPeekable:
    def __init__(self, iterable):
        self.iter = iter(iterable)
        self.suivant = None
        self.fini = False
        self.courant = None
        """ initialisation """

    def __iter__(self):
        return self

    def __next__(self):
        try:
            self.courant = self.suivant
            self.suivant = next(self.iter)

            return self.courant

        except StopIteration:
            self.fini = True
            raise StopIteration
        """ avance l'itérateur et renvoie l'element suivant """

    def peek(self):
        """ renvoie l'element suivant de l'iterateur s'il existe et None sinon """
        return None if self.fini else self.suivant
    
    #def avance(self):
        """ stocker le element suivant de l'iterateur dans self.suivant. """
        self.suivant = next(self.iter)



I = IterateurPeekable(range(4))
for i in I:
    print('valeur courante :', i, ', valeur suivante :', I.peek())