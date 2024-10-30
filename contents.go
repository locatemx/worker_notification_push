package main

import "fmt"

type ContentGenerated struct {
	Priority bool
	Headings NotificationContents `json:"headings"`
	Contents NotificationContents `json:"contents"`
}

func GetContent(event Detail, deviceName string) (ContentGenerated, error) {
	switch event.AlarmCode {
	case "input_1_active":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se ha presionado el boton de SOS",
				En: "The SOS button has been pressed",
			},
		}, nil
	case "ignition_on":
		return ContentGenerated{
			Priority: false,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El vehiculo se ha encendido",
				En: "The vehicle has started",
			},
		}, nil
	case "ignition_off":
		return ContentGenerated{
			Priority: false,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El vehiculo se ha apagado",
				En: "The vehicle has turned off",
			},
		}, nil
	case "external_low":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: fmt.Sprintf("La batería del vehículo está baja %fv", event.Voltages.External),
				En: fmt.Sprintf("Vehicle battery is low %fv", event.Voltages.External),
			},
		}, nil
	case "external_off":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El dispositivo GPS se ha desconectado de la energía",
				En: "GPS device has been disconnected from power",
			},
		}, nil
	case "external_on":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El dispositivo GPS se ha conectado a la energía",
				En: "The GPS device has been connected to power",
			},
		}, nil
	case "internal_low":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: fmt.Sprintf("Batería baja en el dispositivo GPS %fv", event.Voltages.Internal),
				En: fmt.Sprintf("Low battery on GPS device %fv", event.Voltages.Internal),
			},
		}, nil
	case "gps_antenna_off":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se ha desconectado la antena del dispositivo GPS",
				En: "The antenna of the GPS device has been disconnected",
			},
		}, nil
	case "speeding":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: fmt.Sprintf("El vehículo va a exceso de velocidad %dkm/h", int(event.Gps.Speed)),
				En: fmt.Sprintf("The vehicle is speeding %dkm/h", int(event.Gps.Speed)),
			},
		}, nil
	case "sleep":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El dispositivo GPS ha entrado en modo ahorro de batería",
				En: "The GPS device has entered battery saving mode",
			},
		}, nil
	case "wakeup":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El dispositivo GPS ha salido del modo ahorro de batería",
				En: "The GPS device has exited battery saving mode",
			},
		}, nil
	case "tow":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se detectó que el vehículo está siendo remolcado",
				En: "It was detected that the vehicle is being towed",
			},
		}, nil
	case "harsh_accelerate":
		return ContentGenerated{
			Priority: false,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se detectó que el vehículo realizó una fuerte aceleración",
				En: "It was detected that the vehicle made a harsh acceleration",
			},
		}, nil
	case "harsh_braking":
		return ContentGenerated{
			Priority: false,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se detectó que el vehículo realizó un frenada brusca",
				En: "It was detected that the vehicle made a harsh braking",
			},
		}, nil
	case "harsh_turning":
		return ContentGenerated{
			Priority: false,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se detectó que el vehículo realizó un giro brusco",
				En: "It was detected that the vehicle made a harsh turning",
			},
		}, nil
	case "impact":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "Se ha detectado un posible impacto del vehículo",
				En: "A possible vehicle impact has been detected",
			},
		}, nil
	case "dismantling":
		return ContentGenerated{
			Priority: true,
			Headings: NotificationContents{
				Es: fmt.Sprintf("Alarma en %s", deviceName),
				En: fmt.Sprintf("Alarm in %s", deviceName),
			},
			Contents: NotificationContents{
				Es: "El dispositivo GPS fue removido",
				En: "GPS device was removed",
			},
		}, nil
	default:
		return ContentGenerated{}, fmt.Errorf("no se encontro el contenido")
	}
}
