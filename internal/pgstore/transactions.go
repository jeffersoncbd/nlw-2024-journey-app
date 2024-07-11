package pgstore

import (
	"context"
	"fmt"
	"journey/internal/api/spec"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (q *Queries) CreateTrip(ctx context.Context, pool *pgxpool.Pool, params spec.CreateTripRequest) (uuid.UUID , error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
        return uuid.Nil, fmt.Errorf("pgstore: Falha ao tentar criar transação para CreateTrip: %w", err)
    }

	defer tx.Rollback(ctx)

	qtx := q.WithTx(tx)

	tripID, err := qtx.InsertTrip(ctx, InsertTripParams{
		Destination: params.Destination,
		OwnerEmail:  string(params.OwnerEmail),
		OwnerName:   params.OwnerName,
		StartsAt:    pgtype.Timestamp{Valid:true, Time:params.StartsAt},
		EndsAt:      pgtype.Timestamp{Valid:true, Time:params.EndsAt},
	})
	if err != nil {
        return uuid.Nil, fmt.Errorf("pgstore: Falha ao tentar inserir viagem para CreateTrip: %w", err)
    }

	participants := make([]InviteParticipantsToTripParams, len(params.EmailsToInvite))
	for i, email := range params.EmailsToInvite {
		participants[i] = InviteParticipantsToTripParams{
            TripID: tripID,
            Email:  string(email),
        }
	}

	if _, err := qtx.InviteParticipantsToTrip(ctx, participants); err != nil {
		return uuid.Nil, fmt.Errorf("pgstore: Falha ao tentar convidar participantes para CreateTrip: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("pgstore: Falha ao tentar confirmar transação para CreateTrip: %w", err)
	}

	return tripID, nil
}