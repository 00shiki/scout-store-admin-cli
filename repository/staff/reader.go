package staff

import (
	"fmt"
	"scout-store-admin-cli/entity"
)

func (r *Repository) FetchStaffs() ([]entity.Staff, error) {
	query := `
			SELECT s.StaffID, s.StaffName, s.Email, s.RoleID, r.RoleName
			FROM Staff s
					 JOIN Roles r
						  ON s.RoleID = r.RoleID`
	var staffs []entity.Staff
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying staffs: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var staff entity.Staff
		err := rows.Scan(&staff.ID, &staff.Name, &staff.Email, &staff.Role.ID, &staff.Role.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning staff: %v", err)
		}
		staffs = append(staffs, staff)
	}
	return staffs, nil
}

func (r *Repository) FetchStaffByID(staffID int) (entity.Staff, error) {
	query := `
			SELECT s.StaffID, s.StaffName, s.Email, s.RoleID, r.RoleName
			FROM Staff s
					 JOIN Roles r
						  ON s.RoleID = r.RoleID
			WHERE s.StaffID = ?`
	var staff entity.Staff
	rows, err := r.DB.Query(query, staffID)
	if err != nil {
		return entity.Staff{}, fmt.Errorf("error querying staff: %v", err)
	}
	defer rows.Close()
	for !rows.Next() {
		return entity.Staff{}, fmt.Errorf("staff not found")
	}
	err = rows.Scan(&staff.ID, &staff.Name, &staff.Email, &staff.Role.ID, &staff.Role.Name)
	if err != nil {
		return entity.Staff{}, fmt.Errorf("error scanning staff: %v", err)
	}
	return staff, nil
}
