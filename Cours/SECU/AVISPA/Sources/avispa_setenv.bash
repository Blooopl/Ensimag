# Pour configurer avispa sur ensipc si probleme avec AVISPA_PACKAGE : enregistrer ce 
#    fichier dans avispa_setenv.bash et taper la commande :
# source avispa_setenv.bash
_avispa_dir=~/.avispa
mkdir -p $_avispa_dir/logs
ln -sf /opt/avispa-1.1/bin $_avispa_dir/
export AVISPA_PACKAGE="$_avispa_dir"
