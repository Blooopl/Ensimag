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
            self.intervalles = [range(0,nombre_ressources)]
        


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
        pass

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
    #test()

    res = Ressources(20, [[0,2],[4,9],[11,17]])
    print('debut: res =', str(res))
    R1 = res.reserve(3)
    print("R1 :", str(R1))


