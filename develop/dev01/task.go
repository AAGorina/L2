package dev01

import (
	"github.com/beevik/ntp"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func CurrentTime() (time.Time, error) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	time := time.Now().Add(response.ClockOffset)
	return time, err
}
