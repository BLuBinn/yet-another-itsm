INSERT INTO business_units (id, domain_name, tenant_id, name) VALUES 
    ('550e8400-e29b-41d4-a716-446655440019','masangroupcorp.onmicrosoft.com' ,'3ad392ec-0751-4a55-9d34-66fe57b061a4', 'Masangroup'),
    ('550e8400-e29b-41d4-a716-446655440020','WinMart.Masangroup.com' ,'3ad392ec-0751-4a55-9d34-66fe57b061a4', 'Phòng kinh doanh'),
    ('550e8400-e29b-41d4-a716-446655440021','phuclong.masangroup.com' ,'3ad392ec-0751-4a55-9d34-66fe57b061a4', 'Phòng nhân sự'),
    ('550e8400-e29b-41d4-a716-446655440022','crownx.masangroup.com' ,'3ad392ec-0751-4a55-9d34-66fe57b061a4', 'Phòng công nghệ thông tin'),
    ('550e8400-e29b-41d4-a716-446655440023','wineco.masangroup.com' ,'3ad392ec-0751-4a55-9d34-66fe57b061a4', 'Phòng tài chính');
INSERT INTO departments (id, business_unit_id, name) VALUES
    ('550e8400-e29b-41d4-a716-446655440029', '550e8400-e29b-41d4-a716-446655440019', 'Admin Department'),
    ('550e8400-e29b-41d4-a716-446655440030', '550e8400-e29b-41d4-a716-446655440020', 'Sales Team A'),
    ('550e8400-e29b-41d4-a716-446655440031', '550e8400-e29b-41d4-a716-446655440020', 'Sales Team B'),
    ('550e8400-e29b-41d4-a716-446655440032', '550e8400-e29b-41d4-a716-446655440021', 'HR Operations'),
    ('550e8400-e29b-41d4-a716-446655440033', '550e8400-e29b-41d4-a716-446655440022', 'IT Support'),
    ('550e8400-e29b-41d4-a716-446655440034', '550e8400-e29b-41d4-a716-446655440023', 'Finance Team');
INSERT INTO users (id, home_tenant_id ,azure_ad_object_id, department_id, mail, display_name, given_name, sur_name, job_title, office_location) VALUES 
    ('550e8400-e29b-41d4-a716-446655440010', '3ad392ec-0751-4a55-9d34-66fe57b061a4', '550e8400-e29b-41d4-a716-446655440111', '550e8400-e29b-41d4-a716-446655440029', 'anh.admin@masangroupcorp.onmicrosoft.com', 'Admin Đỗ Đức Anh', 'Admin Đỗ Đức', 'Anh', 'Administrator', '23 Lê Duẩn - Quận 1 - Hồ Chí Minh'),
    ('550e8400-e29b-41d4-a716-446655440011', '3ad392ec-0751-4a55-9d34-66fe57b061a4', '550e8400-e29b-41d4-a716-446655440111', '550e8400-e29b-41d4-a716-446655440030', 'anhdt@WinMart.Masangroup.com', 'Dương Thành Anh', 'Dương Thành', 'Anh', 'Sales Manager', '23 Lê Duẩn - Quận 1 - Hồ Chí Minh'),
    ('550e8400-e29b-41d4-a716-446655440012', '3ad392ec-0751-4a55-9d34-66fe57b061a4', '550e8400-e29b-41d4-a716-446655440111', '550e8400-e29b-41d4-a716-446655440032', 'datdt@phuclong.masangroup.com', 'Đỗ Thành Đạt', 'Đỗ Thành', 'Đạt', 'HR Manager', '23 Lê Duẩn - Quận 1 - Hồ Chí Minh'),
    ('550e8400-e29b-41d4-a716-446655440013', '3ad392ec-0751-4a55-9d34-66fe57b061a4', '550e8400-e29b-41d4-a716-446655440111', '550e8400-e29b-41d4-a716-446655440030', 'duyenttm@WinMart.Masangroup.com', 'Trần Thị Mỹ Duyên', 'Trần Thị Mỹ', 'Duyên', 'Sales Executive', '23 Lê Duẩn - Quận 1 - Hồ Chí Minh'),
    ('550e8400-e29b-41d4-a716-446655440014', '3ad392ec-0751-4a55-9d34-66fe57b061a4', '550e8400-e29b-41d4-a716-446655440111', '550e8400-e29b-41d4-a716-446655440032', 'dungtt@phuclong.masangroup.com', 'Dương Thị Thùy Dung', 'Dương Thị Thùy', 'Dung', 'HR Specialist', '23 Lê Duẩn - Quận 1 - Hồ Chí Minh');
