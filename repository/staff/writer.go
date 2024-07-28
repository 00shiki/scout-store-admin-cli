package staff

import (
	"fmt"
	"scout-store-admin-cli/entity"
)

func (r *Repository) AddStaff(staff entity.Staff) error {
	query := `INSERT INTO Staff (StaffName, RoleID, Email) VALUES (?, ?, ?)`
	_, err := r.DB.Exec(query, staff.Name, staff.Role.ID, staff.Email)
	if err != nil {
		return fmt.Errorf("could not insert staff: %w", err)
	}
	return nil
}
