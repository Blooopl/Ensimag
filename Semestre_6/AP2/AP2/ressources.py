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
        """

        ecarts_dispo = self.ecarts_disponible()
        somme_ecarts_dispo = 0
        for ecart in ecarts_dispo:
            somme_ecarts_dispo+=ecart[0]
        if somme_ecarts_dispo < ressources_demandees:
            return None

        liste_allocation = []
        reste_a_allouer = ressources_demandees
        while reste_a_allouer != 0 :

            ecarts_dispo = self.ecarts_disponible()
            print(ecarts_dispo)
            
            taille_optimum = ecarts_dispo[1][1]
            id_optimum = 0
            for ecart in ecarts_dispo:

                valeur_ecart = ecart[0]
                if ressources_demandees - valeur_ecart == 0:
                    taille_optimum = valeur_ecart
                    id_optimum = ecart[1]

                elif ressources_demandees - valeur_ecart > 0:
                    if valeur_ecart > taille_optimum:
                        taille_optimum = valeur_ecart
                        id_optimum = ecart[1]
                
                else:
                    if valeur_ecart < taille_optimum:
                        taille_optimum = valeur_ecart
                        id_optimum = ecart[1]
            
            reste_a_allouer = reste_a_allouer - taille_optimum

            debut_allocation = self.intervalles[id_optimum-1][1]
            fin_allocation =self.intervalles[id_optimum-1][1] + id_optimum
            
            
            liste_allocation.append([debut_allocation,fin_allocation])

            self.intervalles.append([debut_allocation,fin_allocation])

        print(liste_allocation)
        res = Ressources(self.nombre_ressources)
        for allocation in liste_allocation:
            res.intervalles.append(allocation)

        
        return res
    def ecarts_disponible(self):
        """
        Renvoie une liste de l'écart entre chaque intervalle
        """
        liste = []
        
        liste.append([self.intervalles[0][0],0])
        
        for id_intervalle in range(0,len(self.intervalles)-1):
            fin_premier_intervalle = self.intervalles[id_intervalle][1]
            debut_second_intervalle = self.intervalles[id_intervalle+1][0]
            liste.append([debut_second_intervalle - fin_premier_intervalle,id_intervalle+1])
        
        liste.append([self.nombre_ressources-self.intervalles[-1][1],len(self.intervalles)])
        return liste


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
    a = Ressources(10, [[0,6],[6,10]])
    """
    print(Ressources(10,[[0, 6], [6, 10]]).verification_invariants())
    print(Ressources(10,[[0, 7], [6, 10]]).verification_invariants())
    print(Ressources(10,[[1, 6], [6, 10] , [0,1]]).verification_invariants())
    print(Ressources(10,[[1, 6], [6, 11]]).verification_invariants())"""

    """print(Ressources(10,[[0,5],[6,10]]).disponible(6))
    print(Ressources(10,[[0,5],[6,10]]).disponible(5))
    print(Ressources(10,[[1,6],[6,8]]).disponible(8))


    print(Ressources(15,[[0,4],[7,10]]).ecarts_disponible())"""

    res = Ressources(10, [[0, 4]])
    print('debut: res =', str(res))
    R1 = res.reserve(1)
    print(R1)
