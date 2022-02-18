package formator

import "github.com/qiangyt/loggen/pkg/config"

type Formator interface {
	Format(state config.State) string
}
