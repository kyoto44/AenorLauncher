<div align="center">
	<h1>Aenor Launcher</h1>
    <h4 align="center">
	   Простой и быстрый лаунчер для игры Северный Клинок
	</h4>
</div>

## 🎯 Основные возможности

* 🔑Автологин
* 👑Ник и гильдия в статусе Discord (Список поддерживаемых гильдий ниже)
* 🔄Поддержка обновлений клиента игры
* 💼Сохранение игровых настроек после обновлений
* 🚀Многопоточная загрузка ресурсов
* 🔒Проверка целостности игровых файлов
* 🐧Поддержка Linux

## ⬇️ Скачать 
[Версия для Windows](https://github.com/kyoto44/AenorLauncher/releases/download/1.0/AenorLauncher1.0-win.zip) ![GitHub Releases (by Asset)](https://img.shields.io/github/downloads/kyoto44/AenorLauncher/1.0/AenorLauncher1.0-win.zip?style=flat-square)

[Версия для Linux](https://github.com/kyoto44/AenorLauncher/releases/download/1.0/AenorLauncher1.0-linux.zip) ![GitHub Releases (by Asset)](https://img.shields.io/github/downloads/kyoto44/AenorLauncher/1.0/AenorLauncher1.0-linux.zip?style=flat-square)

## 👋 Как использовать

### Windows 
Поместить AenorLauncher.exe и config.txt в одну папку и запустить AenorLauncher.exe

### Linux
Для запуска игры необходимо установить: 
* Wine
* DXVK
* D3DX9_43.dll из Winetricks

И в Терминале: 
```
chmod +x AenorLauncher && ./AenorLauncher
```

## ⚙️ Настройки
Пример файла config.txt 
```
username = "alexxxx" # Логин
password = "alexxxx" # Пароль
nickname = "Персонаж" # Ник
guild = "Гильдия" # Гильдия
```

## 👑 Гильдии
Поддерживаемые гильдии на данный момент (с логотипом): 
* КАЭР МОРХЕН
* ИМПЕРИЯ
* FORCE
* НЕМЕЗИДА

## Лицензия

AenorLauncher is provided under the [MIT License](https://github.com/kyoto44/AenorLauncher/blob/master/LICENSE) © Dmitry Gomzyakov.