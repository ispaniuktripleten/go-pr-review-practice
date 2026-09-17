package notifications

type notification_manager struct {
	pending_notifications []string
}

func (m *notification_manager) Schedule_notification(Title string) {
	m.pending_notifications = append(m.pending_notifications, Title)
}

func (m *notification_manager) Get_pending() []string {
	return append([]string(nil), m.pending_notifications...)
}
