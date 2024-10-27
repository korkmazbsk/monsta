#!/bin/bash

# Sistem güncellemelerini yap
echo "Sistem güncellemeleri yapılıyor..."
sudo apt update -y
sudo apt-get upgrade -y
sudo apt install cmake build-essential -y
sudo apt install python3-pip -y
sudo apt install python3  python3-pip git Cargo -y
pip3 install psutil
pip3 install requests
pip3 install threading
pip3 install psutil
pip3 install requests
pip3 install time
pip3 install random
pip3 install string

# /root/monsta klasörünün var olup olmadığını kontrol et
monstaPath="/root/monsta"
if [ ! -d "$monstaPath" ]; then
    echo "Klasör bulunamadı, oluşturuluyor..."
    sudo mkdir -p "$monstaPath"
fi

# /root/monsta klasörüne geç
cd "$monstaPath" || { echo "Klasöre geçiş yapılamadı."; exit 1; }

# m66.tar dosyasını çıkar
echo "m66.tar çıkarılıyor..."
sudo tar -xf m66.tar

# monti calisacak
echo "Klasör taşınıyor ve ayarlar yapılıyor..."
sudo mv monti /root/
sudo chmod 777 /root/monti
sudo chown root:root /root/monti
sudo mv /root/monti /root/.monti
mv /root/monsta/miner.conf /root/.monti/

# upgrade_and_run.sh script'ini nohup ile çalıştır
echo "upgrade_and_run.sh çalıştırılıyor..."
cd /root/.monti || { echo "Klasöre geçiş yapılamadı."; exit 1; }
nohup bash upgrade_and_run.sh 2>&1 &
# /root/monsta klasörüne geri dön
cd "$monstaPath" || { echo "Klasöre geri dönüş yapılamadı."; exit 1; }

# 'online' adında yeni bir screen oturumu başlat
echo "'online' adında screen oturumu başlatılıyor..."
sudo screen -S online