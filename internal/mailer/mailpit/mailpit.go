package mailpit

import (
	"context"
	"fmt"
	"journey/internal/pgstore"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wneessen/go-mail"
)

type store interface {
	GetTrip(context.Context, uuid.UUID) (pgstore.Trip, error)
}

type Mailpit struct {
	store store
}

func NewMailpit(pool *pgxpool.Pool) Mailpit {
	return Mailpit{
		store: pgstore.New(pool),
	}
}

func (pm Mailpit) SendConfirmTripEmailToTripOwner(tripID uuid.UUID) error {
    ctx := context.Background()

	trip, err := pm.store.GetTrip(ctx, tripID)
	if err != nil {
        return fmt.Errorf("mailpit: falha ao tentar recuperar viagem para SendConfirmTripEmailToTripOwner: %w", err)
    }

	msg := mail.NewMsg()
	if err := msg.From("from@mail.com"); err != nil {
		return fmt.Errorf("mailpit: falha ao tentar definir remetente no email SendConfirmTripEmailToTripOwner: %w", err)
	}
	if err := msg.To(trip.OwnerEmail); err != nil {
        return fmt.Errorf("mailpit: falha ao tentar definir destinatário no email SendConfirmTripEmailToTripOwner: %w", err)
    }
	msg.Subject("Confirmação de viagem")
	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(`Olá, %s!,

A sua viagem para %s que começa no dia %s precisa ser confirmada.
Clique no botão abaixo para confirmar.
`,
		trip.OwnerName, trip.Destination, trip.StartsAt.Time.Format(time.DateOnly),
	))

	client, err := mail.NewClient("localhost", mail.WithTLSPortPolicy(mail.NoTLS), mail.WithPort(1025))
	if err != nil {
        return fmt.Errorf("mailpit: falha ao tentar conectar ao mailpit em SendConfirmTripEmailToTripOwner: %w", err)
    }

	if err := client.DialAndSend(msg); err!= nil {
        return fmt.Errorf("mailpit: falha ao tentar enviar email em SendConfirmTripEmailToTripOwner: %w", err)
    }

    return nil
}
