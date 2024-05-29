package push

//func New(cfg []*client.Config) *Manager {
//	m := &Manager{
//		m: make(map[string]client.Client),
//	}
//	for _, v := range cfg {
//		key := v.GetKey()
//		if m.m[key] == nil {
//			if c, err := v.Dial(); err == nil {
//				m.m[key] = c
//			}
//		}
//	}
//	logs.Debug(m)
//	return m
//}
//
//type Manager struct {
//	m map[string]client.Client
//}
//
//func (this *Manager) Publish(key string, topic string, data interface{}) error {
//
//	c, ok := this.m[key]
//	if !ok {
//		return fmt.Errorf("客户端不存在: %s", key)
//	}
//
//	err := c.Publish(topic, data)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
