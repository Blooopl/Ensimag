"""
manipulations complexes de tableaux : listes d'intervalles.
"""

from bisect import bisect



def fusion(inter1, inter2):
    # merge intervals and sort them by their left boundaries
    pass
        

class Ressources:
    """
    On stocke une liste de ressources, compressee par plages contigues.
    """

    def __init__(self, nombre_ressources, intervalles=None):
        self.nombre_ressources = nombre_ressources
        if intervalles is not None:
            self.intervalles = intervalles
        else:
            self.intervalles = [[0,nombre_ressources]]
        


    def verification_invariants(self):
        """
        pas d'intersections entre les intervalles,
        les intervalles sont tries du plus petit au plus gros indice
        """
        max_actuel = 0

        for intervalle in self.intervalles:
            debut = intervalle[0]
            fin = intervalle[1]

            if debut >= max_actuel and fin <= self.nombre_ressources:
                max_actuel = fin
            else:
                return False

        return True

    def disponible(self, indice):
        """
        renvoie si l'indice donne est disponible dans la ressource.
        """
        for intervalle in self.intervalles:
            debut = intervalle[0]
            fin = intervalle[1]
            if indice < fin and indice >= debut:
                return True
            
            if debut > indice:
                return False
        
        return False

    def reserve(self, ressources_demandees):
        """
        enleve le nombre de ressources demandees.
        renvoie les ressources correspondant aux plages reservees.


        On essaye d'allouer les ressources en un minimum de nombres de plages
        De + on essaye lorsque c'est possible de ne pas utiliser des plages libre plus grandes que nécessaires
        C'est à dire que si on veut alouer 3 ressources alors que self.intervalles = [[0,4],[5,10]], on va éviter de "gacher" la deuxième plages qui pourrait être utile si on veut de nouveaux allouer un truc de grandes tailles
        """
        
        res = Ressources(self.nombre_ressources,[])
        reste_a_allouer = ressources_demandees

        liste_allocation = []
        while reste_a_allouer != 0 :
            
            alloc_opti = [0,0]
            taille_optimum = None

            for id_intervalle in range(len(self.intervalles)):

                intervalle = self.intervalles[id_intervalle]

                debut_intervalle = intervalle[0]
                fin_intervalle = intervalle[1]

                dispo = fin_intervalle - debut_intervalle

                if dispo == reste_a_allouer:
                    alloc_opti = [debut_intervalle,fin_intervalle-1]     
                    taille_optimum = reste_a_allouer

                if dispo > reste_a_allouer:

                    if taille_optimum == None or (dispo < taille_optimum or taille_optimum < reste_a_allouer):
                        taille_optimum = dispo
                        alloc_opti = [debut_intervalle,debut_intervalle+reste_a_allouer-1]
 
                if dispo < reste_a_allouer:
                    
                    if taille_optimum == None or dispo > taille_optimum:
                        taille_optimum = dispo
                        alloc_opti = [debut_intervalle,fin_intervalle-1]

            if taille_optimum > reste_a_allouer:
                taille_optimum = reste_a_allouer

            reste_a_allouer -= taille_optimum

            liste_allocation.append(alloc_opti)

            for id_intervalle in range(len(self.intervalles)):

                intervalle = self.intervalles[id_intervalle]
                debut_intervalle = intervalle[0]
                fin_intervalle = intervalle[1]

                if debut_intervalle == alloc_opti[0]:
                    
                    if fin_intervalle-1 == alloc_opti[1]:
                        self.intervalles.pop(id_intervalle)
                        break
                    else:
                        self.intervalles[id_intervalle] = [alloc_opti[1]+1,fin_intervalle]

        for allocation in liste_allocation:
            res.intervalles.append([allocation[0],allocation[1]+1])

        
        return res

    def retourne(self, ressources_rendues):
        """
        remet les plages de ressources donnees dans le systeme.
        """
        self.intervalles = fusion(self.intervalles,ressources_rendues.intervalles)

        self.simplification_intervalles()

    def simplification_intervalles(self):

        id_intervalle = 0
        while id_intervalle < len(self.intervalles) - 1:

            while id_intervalle < len(self.intervalles) - 1  and self.intervalles[id_intervalle][1] == self.intervalles[id_intervalle+1][0] :
                print("oui")
                self.intervalles[id_intervalle] = [self.intervalles[id_intervalle][0],self.intervalles[id_intervalle+1][1]]
                self.intervalles.pop(id_intervalle+1)
            id_intervalle +=1




    def __str__(self):
        """
        renvoie une chaine 'visuelle' des ressources libres/utilisees.
        par exemple, '|x..xxxxx...|' indique qu'il y a 10 ressources,
        les ressources 0, 3-7 sont libres.
        """
        liste_dispo = ["."]*self.nombre_ressources
        for intervalle in self.intervalles:
            debut = intervalle[0]
            fin = intervalle[1]
            for i in range(debut,fin):
                liste_dispo[i] = "x"

        chaine = "".join(liste_dispo)

        return chaine

    def trier(self):
        
        changement = 1
        while changement != 0:
            changement = 0

            for id_intervalle in range(len(self.intervalles)-1):
                if self.intervalles[id_intervalle][0] > self.intervalles[id_intervalle+1][0]:
                    self.intervalles[id_intervalle+1],self.intervalles[id_intervalle] = self.intervalles[id_intervalle],self.intervalles[id_intervalle+1]
                    changement += 1

 
       
def fusion(inter1,inter2):

    liste = []

    liste_1 = inter1.copy()
    liste_2 = inter2.copy()

    while len(liste_1) != 0 and len(liste_2) != 0:
        if liste_1[0][0] < liste_2[0][0]:
            liste.append(liste_1[0])
            liste_1.pop(0)
        else:
            liste.append(liste_2[0])
            liste_2.pop(0)
    
    return liste+liste_1+liste_2

def test():
    """
    on teste en gruyerisant une ressource.
    """
    ressources = Ressources(10)
    print("Disponibles :", ressources)
    print("on commence par tout reserver, il ne reste donc plus rien:")
    reservees = [ressources.reserve(c) for c in (2, 2, 3, 2, 1)]
    print('reservees = ', [str(r) for r in reservees])
    print("Disponibles :", ressources)
    print("on rend deux ressources:")
    ressources.retourne(reservees[1])
    print("Disponibles :", ressources)
    print("on rend encore deux ressources, mais plus loin:")
    ressources.retourne(reservees[3])
    print("Disponibles :", ressources)
    print("on reserve trois ressources sur les quatres disponibles:")
    print("Reservees   :", ressources.reserve(3))
    print("Disponibles :", ressources)
    print("Les intervalles :", ressources.intervalles)


if __name__ == "__main__":
    test()
    """
    res = Ressources(20)
    res.trier()
    print(res.intervalles)
    print('debut: res =', str(res))
    R1 = res.reserve(3)
    print("R1 :", str(R1))

    print(fusion([[0, 1], [5, 6]], [[1, 2], [4, 5]]))
    fusion([[0, 1], [5, 6]], [[1, 2], [4, 5]])

    res = Ressources(20, [[1,2],[2,3],[3,4],[4,5],[7,10],[13,15],[15,16],[17,18],[18,19]])
    res.simplification_intervalles()
    print(res.intervalles)"""
