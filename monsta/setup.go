package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
)

func runCommand(command string, args ...string) {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Error running command %s: %v\n", command, err)
    }
}

func main() {
    // 1. Sistem güncellemelerini yap
    fmt.Println("Sistem güncellemeleri yapılıyor...")
    runCommand("sudo", "apt", "update", "-y")
    runCommand("sudo", "apt-get", "upgrade", "-y")

    // 2. /root/monsta klasörünün var olup olmadığını kontrol et
    monstaPath := "/root/monsta"
    if _, err := os.Stat(monstaPath); os.IsNotExist(err) {
        fmt.Println("Klasör bulunamadı, oluşturuluyor...")
        runCommand("sudo", "mkdir", "-p", monstaPath)
    }

    // 3. /root/monsta klasörüne geç
    err := os.Chdir(monstaPath)
    if err != nil {
        log.Fatalf("Klasöre geçiş yapılamadı: %v\n", err)
    }

    // 4. m66.tar dosyasını çıkar
    fmt.Println("m66.tar çıkarılıyor...")
    runCommand("sudo", "tar", "-xf", "m66.tar")

    // 5. monti klasörünü /root altına gizli olarak taşı ve ayarları yap
    fmt.Println("Klasör taşınıyor ve ayarlar yapılıyor...")
    runCommand("sudo", "mv", "monti", "/root/")
    runCommand("sudo", "chmod", "700", "/root/monti")
    runCommand("sudo", "chown", "root:root", "/root/monti")
    runCommand("sudo", "mv", "/root/monti", "/root/.monti")

    // 6. upgrade_and_run.sh script'ini screen ile çalıştır
    fmt.Println("upgrade_and_run.sh çalıştırılıyor...")
    runCommand("sudo", "screen", "-dmS", "monstar", "bash", "-c", "cd /root/.monti && nohup bash upgrade_and_run.sh > /dev/null 2>&1")

    // 7. /root/monsta klasörüne geri dön
    err = os.Chdir(monstaPath)
    if err != nil {
        log.Fatalf("Klasöre geri dönüş yapılamadı: %v\n", err)
    }

    // 8. 'montim' adında yeni bir screen oturumu başlat
    fmt.Println("'montim' adında screen oturumu başlatılıyor...")
    runCommand("sudo", "screen", "-dmS", "montim")

    fmt.Println("Tüm işlemler başarıyla tamamlandı.")
}
