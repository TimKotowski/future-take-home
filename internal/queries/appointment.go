package queries

const (
	InsertAppointmentQuery = `
	INSERT INTO appointments (
		id,  
		trainer_id, 
		user_id, 
		start_slot,
		end_slot,
		status
	)
	VALUES(
		$1, 
		$2,
		$3, 
		$4, 
		$5,
		cast($6 as appointments_status)
	) RETURNING id, trainer_id, user_id, start_slot, end_slot, status, created_at, updated_at
`
	GetAppointmentByTrainerQuery = `
		SELECT 
			a.id,
			a.trainer_id,
			a.user_id,
			a.start_slot,
			a.end_slot,
			a.status,
			a.created_at,
			a.updated_at
		FROM appointments a 
			INNER JOIN trainers t on t.id = a.trainer_id
				where a.trainer_id = $1
				and a.status = $2
`

	GetAppointmentByTrainerAndStartSlotAndEndSlotQuery = `
		SELECT 
			a.id,
			a.trainer_id,
			a.user_id,
			a.start_slot,
			a.end_slot,
			a.status,
			a.created_at,
			a.updated_at
		FROM appointments a 
			INNER JOIN trainers t on t.id = a.trainer_id
				where a.trainer_id = $1
				and a.start_slot >= $2 and a.end_slot <= $3
				and a.status = $4
`
)
