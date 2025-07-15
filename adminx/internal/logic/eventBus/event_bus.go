/**
 * @Company:
 * @Author: yxf
 * @Description:
 * @Date: 2024/1/25 16:22
 */

package eventBus

import (
	"github.com/asaskevich/EventBus"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterEventBus(EventBus.New())
}
