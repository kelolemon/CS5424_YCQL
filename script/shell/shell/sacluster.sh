ssh x0 "./kill.sh"
ssh x1 "./kill.sh"
ssh x2 "./kill.sh"
ssh x3 "./kill.sh"
ssh x4 "./kill.sh"

ssh x0 "./start_master.sh xcnd5.comp.nus.edu.sg" && \
ssh x1 "./start_master.sh xcnd6.comp.nus.edu.sg" && \
ssh x3 "./start_master.sh xcnd8.comp.nus.edu.sg" && 

ssh x2 "./start_tserver.sh xcnd7.comp.nus.edu.sg" && \
ssh x4 "./start_tserver.sh xcnd50.comp.nus.edu.sg"

