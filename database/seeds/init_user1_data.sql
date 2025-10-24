-- =====================================================
-- Insert test users
-- =====================================================
INSERT INTO `user` (`username`, `email`, `password`, `nickname`, `status`, `created_at`, `updated_at`)
VALUES ('testuser1', 'test1@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', '测试用户1',
        1, NOW(), NOW()),
       ('testuser2', 'test2@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', '测试用户2',
        1, NOW(), NOW()),
       ('testuser3', 'test3@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', '测试用户3',
        1, NOW(), NOW());