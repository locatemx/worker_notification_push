package main

type AccountDevice struct {
	Alias   string `json:"alias"`
	Account string `json:"account"`
}

func StringPtr(s string) *string {
	return &s
}

func NotifyAlarm(detail Detail, accounts []AccountDevice) (err error) {
	for _, account := range accounts {
		//? OBTENER CONTENIDO BASADO EN
		content, err := GetContent(detail, account.Alias)
		if err != nil {
			return err
		}

		//if content.Priority {
		//! ENVIAR NOTIFICATION PRIORITARIA
		//}

		//! ENVIO DEPENDIENDO FILTER
		//? if Notification.Filter == "account"
		// err = OneSignalSendNotification(NotificationBody{
		// 	Filters: []NotificationFilter{
		// 		{
		// 			Field:    StringPtr("tag"),
		// 			Key:      StringPtr("account_id"),
		// 			Relation: StringPtr("="),
		// 			Value:    StringPtr(Notification.FilterId),
		// 		},
		// 	},
		// 	Headings:Notification.Content.Headings,
		// 	Contents: Notification.Content.Contents,
		// 	Buttons:  []NotificationButton{},
		// })

		//? if Notification.Filter == "user"
		// err = OneSignalSendNotification(NotificationBody{
		// 	Filters: []NotificationFilter{
		// 		{
		// 			Field:    StringPtr("tag"),
		// 			Key:      StringPtr("user_id"),
		// 			Relation: StringPtr("="),
		// 			Value:    StringPtr(Notification.FilterId),
		// 		},
		// 	},
		// 	Headings:Notification.Content.Headings,
		// 	Contents: Notification.Content.Contents,
		// 	Buttons:  []NotificationButton{},
		// })

		err = OneSignalSendNotification(NotificationBody{
			Filters: []NotificationFilter{
				{
					Field:    StringPtr("tag"),
					Key:      StringPtr("account_id"),
					Relation: StringPtr("="),
					Value:    StringPtr(account.Account),
				},
			},
			Headings: content.Headings,
			Contents: content.Contents,
			Buttons:  []NotificationButton{},
		})

		if err != nil {
			return err
		}
	}

	return nil
}
