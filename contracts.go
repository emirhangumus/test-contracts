package contracts

const (
	WorkflowTaskQueue = "workflow-task-queue"
	ServiceATaskQueue = "service-a-task-queue"
	ServiceBTaskQueue = "service-b-task-queue"

	WorkflowCode = "order.fulfillment.workflow"

	ReserveInventoryActivity = "inventory.reserve"
	ReleaseInventoryActivity = "inventory.release"
	CapturePaymentActivity   = "payment.capture"

	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusReleased  = "released"
	StatusFailed    = "failed"

	ServiceAActivity = ReserveInventoryActivity
	ServiceBActivity = CapturePaymentActivity
)

type OrderItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type OrderWorkflowInput struct {
	OrderID        string      `json:"order_id"`
	CustomerID     string      `json:"customer_id"`
	Items          []OrderItem `json:"items"`
	AmountCents    int64       `json:"amount_cents"`
	IdempotencyKey string      `json:"idempotency_key"`
}

type OrderWorkflowResult struct {
	OrderID     string                     `json:"order_id"`
	Status      string                     `json:"status"`
	Inventory   InventoryReservationResult `json:"inventory"`
	Payment     PaymentCaptureResult       `json:"payment"`
	Compensated bool                       `json:"compensated"`
}

type InventoryReservationRequest struct {
	OrderID        string      `json:"order_id"`
	Items          []OrderItem `json:"items"`
	IdempotencyKey string      `json:"idempotency_key"`
}

type InventoryReservationResult struct {
	OrderID       string      `json:"order_id"`
	ReservationID string      `json:"reservation_id"`
	Status        string      `json:"status"`
	Items         []OrderItem `json:"items"`
}

type InventoryReleaseRequest struct {
	OrderID        string `json:"order_id"`
	ReservationID  string `json:"reservation_id"`
	Reason         string `json:"reason"`
	IdempotencyKey string `json:"idempotency_key"`
}

type InventoryReleaseResult struct {
	OrderID       string `json:"order_id"`
	ReservationID string `json:"reservation_id"`
	Status        string `json:"status"`
}

type PaymentCaptureRequest struct {
	OrderID        string `json:"order_id"`
	CustomerID     string `json:"customer_id"`
	AmountCents    int64  `json:"amount_cents"`
	IdempotencyKey string `json:"idempotency_key"`
}

type PaymentCaptureResult struct {
	OrderID     string `json:"order_id"`
	PaymentID   string `json:"payment_id"`
	Status      string `json:"status"`
	AmountCents int64  `json:"amount_cents"`
}
