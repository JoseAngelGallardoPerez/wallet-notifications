<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\DB;

class InitTables extends Migration
{
    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
    }

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        // skip the migration if there are another migrations
        // It means this migration was already applied
        $migrations = DB::select('SELECT * FROM migrations LIMIT 1');
        if (!empty($migrations)) {
            return;
        }
        $oldMigrationTable = DB::select("SHOW TABLES LIKE 'schema_migrations'");
        if (!empty($oldMigrationTable)) {
            return;
        }

        DB::beginTransaction();

        try {
            app("db")->getPdo()->exec($this->getSql());
        } catch (\Throwable $e) {
            DB::rollBack();
            throw $e;
        }

        DB::commit();
    }

    private function getSql()
    {
        return <<<SQL
            CREATE TABLE `schema_migrations` (
              `version` bigint(20) NOT NULL,
              `dirty` tinyint(1) NOT NULL
            ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

            INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
            (20190822142343, 0);

            CREATE TABLE `settings` (
              `id` int(11) UNSIGNED NOT NULL,
              `name` varchar(255) NOT NULL,
              `value` varchar(255) NOT NULL,
              `description` varchar(255) DEFAULT NULL,
              `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
              `updated_at` timestamp NULL DEFAULT NULL
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

            INSERT INTO `settings` (`id`, `name`, `value`, `description`, `created_at`, `updated_at`) VALUES
            (1, 'email_from', '', 'System e-mail \"From\" address', '2019-10-30 10:46:00', '2019-10-25 06:52:52'),
            (2, 'email_from_name', ' VELMIE_WALLET', 'System e-mail \"From\" name', '2019-10-25 06:52:52', '2019-10-25 06:52:52'),
            (3, 'logo_url', '', 'Logo for signature', '2019-10-30 10:46:09', '2019-10-25 06:52:52'),
            (4, 'mail_signature', 'VELMIE_WALLET team', 'Mail signature', '2019-10-30 10:46:16', '2019-10-25 06:52:52'),
            (5, 'smtp_status', 'enabled', 'SMTP Settings (recommended)', '2019-10-25 06:52:52', '2019-10-25 06:52:52'),
            (6, 'smtp_host', '', 'SMTP server address', '2019-10-30 10:46:25', '2019-10-25 06:52:52'),
            (7, 'smtp_protocol', 'SSL', 'SMTP TLS/SSL (required)', '2019-10-25 06:52:52', '2019-10-25 06:52:52'),
            (8, 'smtp_username', '', 'SMTP username', '2019-10-30 10:46:29', '2019-10-25 06:52:52'),
            (9, 'smtp_port', '', 'SMTP port', '2019-10-30 10:46:31', '2019-10-25 06:52:52'),
            (10, 'smtp_password', '', 'SMTP password', '2019-10-30 10:46:34', '2019-10-25 06:52:52'),
            (11, 'google_firebase_app_id', '', 'Google Firebase App ID', '2019-10-30 10:46:38', '2019-10-25 07:35:24'),
            (12, 'google_firebase_token', '', 'Google Firebase Token', '2019-10-30 10:46:41', '2019-10-25 07:35:24'),
            (13, 'plivo_auth_id', '', 'Auth ID for Plivo API', '2019-10-30 10:46:43', '2019-10-01 09:18:03'),
            (14, 'plivo_auth_token', '', 'Auth Token for Plivo API', '2019-10-30 10:46:46', '2019-10-01 09:18:03'),
            (15, 'sms_from', '', 'Sender phone number', '2019-10-30 10:46:48', '2019-10-01 09:18:03'),
            (16, 'min_balance', '1', 'Minimum Balance', '2019-10-30 10:47:31', '2018-12-05 14:21:47'),
            (17, 'text_for_sms', 'Please, copy or print this message, since it is only going to be shown once. \\nYour TANs: \\n[Tan]', 'SMS', '2019-10-30 10:47:03', '2018-12-05 14:21:47'),
            (18, 'tan_use_plivo', 'true', 'Use Plivo settings for sending sms with tans', '2019-10-01 09:18:02', '2019-10-01 09:18:03'),
            (19, 'current_balance', '0', '', '2019-10-30 10:47:13', '2018-12-03 10:06:08'),
            (20, 'tans_text_for_sms', 'Please, copy or print this message. \nYour TANs: \n[Tan]', 'TANs text for SMS', '2019-10-01 09:18:02', '2019-10-01 09:18:03'),
            (21, 'plivo_min_balance', '1', 'Minimum allowed balance to send SMS', '2019-10-01 09:18:02', '2019-10-01 09:18:03'),
            (22, 'token', '', '', '2019-01-23 17:02:26', '2019-01-23 17:02:26');

            CREATE TABLE `templates` (
              `id` int(11) UNSIGNED NOT NULL,
              `title` varchar(255) NOT NULL,
              `scope` enum('admin','user') DEFAULT NULL,
              `legend` varchar(255) NOT NULL,
              `subject` varchar(255) NOT NULL,
              `content` text,
              `status` enum('enabled','disabled') DEFAULT 'enabled',
              `is_editable` tinyint(1) NOT NULL DEFAULT '1',
              `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
              `updated_at` timestamp NULL DEFAULT NULL,
              `sort` int(11) DEFAULT '0'
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

            INSERT INTO `templates` (`id`, `title`, `scope`, `legend`, `subject`, `content`, `status`, `is_editable`, `created_at`, `updated_at`, `sort`) VALUES
            (1, 'Registration confirmation', 'user', 'Profile creation via sign up form', 'Account details for [UserName] at [SiteName] (pending admin approval)', '[UserName],\n\nThank you for registering at [SiteName]. Your application for an account is currently pending approval. Once it has been approved, you will receive another e-mail containing information about how to log in, set your password, and other details.\n\n--  [SiteName] team', 'enabled', 0, '2019-11-01 07:22:36', NULL, 996),
            (2, 'Cancelled registration', 'user', 'Profile creation via sign up form', 'Your registration request has been canceled.', 'Your registration request has been canceled by an Administrator. Your Profile did not meet the standard requirements.', 'enabled', 0, '2019-11-01 07:22:36', NULL, 992),
            (3, 'Profile activation', 'user', 'Profile creation via sign up form', 'Your profile has been activated', 'Your [SiteName] Profile status is now Active.', 'enabled', 0, '2019-11-01 07:22:36', NULL, 988),
            (4, 'Incoming transaction', 'user', 'Account alerts', 'Incoming transaction', "You\'ve received an incoming transaction. #[AccountNumber].\n[Logo]", 'enabled', 1, '2019-11-01 07:22:55', NULL, 984),
            (5, 'Username & password change', 'user', 'Account alerts', 'Username or password changed', 'Your [SiteName] username or password has been changed.', 'enabled', 1, '2019-11-01 07:22:36', NULL, 980),
            (6, 'Blocked-failed login attempts', 'user', 'System access', 'Failed Login Attempts', 'Your Profile has been blocked. You have reached the maximum number of failed login attempts. Please try again later.', 'enabled', 1, '2019-11-01 07:22:36', NULL, 976),
            (7, 'New profile welcome', 'user', 'Profile creation via other sources', 'An administrator created an account for you at [SiteName]', '[UserName],\n\nA site administrator at [SiteName] has created an account for you. You may now set your account password by clicking this link or copying and pasting it to your browser:\n\n[SetPasswordOneTimeURL]\n\n[SiteName] team', 'enabled', 0, '2019-11-01 07:22:36', NULL, 972),
            (8, 'Profile activation via other sources', 'user', 'Profile creation via other sources', 'Your profile has been activated', 'Your [SiteName] Profile status is now Active.', 'enabled', 0, '2019-11-01 07:22:36', NULL, 968),
            (9, 'Incoming messages', 'user', 'Message alerts', 'New private message at [SiteName].', "Hi [PrivateMessageRecipient],\n\nThis is an automatic reminder from the site [SiteName]. You have received a new private message from [PrivateMessageAuthor].\n\nTo read your message, follow this link:\n[PrivatemsgMessageURL]\n\nIf you don\'t want to receive these emails again, change your preferences here:\n[PrivateMessageRecipientEditURL]", 'enabled', 1, '2019-11-01 07:22:36', NULL, 964),
            (10, 'News notification', 'user', 'News notification', 'New article published on [SiteName]', 'To read recently published [SiteName] News, log into [SiteName] and select the News tab.', 'enabled', 0, '2019-11-01 07:22:36', NULL, 960),
            (12, 'New registration request', 'admin', '', 'New registration request created', '[FirstName] [LastName] has requested to register on your system. Please log in to your Admin interface to view the request.', 'enabled', 0, '2019-11-01 07:22:36', NULL, 952),
            (13, 'New transfer request', 'admin', '', 'New transfer request created', '[FirstName] [LastName] has requested a transfer. Please log in to your Admin interface to view the request.', 'enabled', 1, '2019-11-01 07:22:36', NULL, 948),
            (14, 'Blocked-failed login attempts (Admin)', 'admin', '', 'User Profile Blocked', '[FirstName] [LastName] has reached the maximum number of failed login attempts. The Profile is now blocked.', 'enabled', 1, '2019-11-01 07:22:36', NULL, 944),
            (15, 'Test Email', 'admin', '', 'A test email (please ignore)', 'If you receive this email, it means that your site is configured with valid SMTP credentials and is sending emails correctly\n\n[SiteName]\n\n[SiteLoginUrl]', 'enabled', 1, '2019-11-01 07:22:36', NULL, 940),
            (16, 'Outgoing transaction', 'user', 'Account alerts', 'Your pending transaction was executed.', 'Your transaction #[TransactionId] has been executed successfully.\n[Logo]', 'enabled', 0, '2019-11-01 07:22:36', NULL, 936),
            (17, 'Resend ID Validation', 'user', 'Verification', 'ID Validation', '<p>Dear [FirstName] [LastName],</p><p>An administrator has issued you with a new ID validation link. Click on the following link to try again. \n\n [VerificationLink] \n\n [SiteName]</p>', 'enabled', 0, '2019-11-01 07:22:36', NULL, 932),
            (18, 'User ID Validation Successful', 'admin', 'Verification', 'User ID Validation successful', '<p>The user [UserName] successful the ID Validation process.</p><p>You can view the results of the ID validation here: \n\n [VerificationLink]</p>', 'enabled', 0, '2019-11-01 07:22:36', NULL, 928),
            (19, 'User ID Validation Failed', 'admin', 'Verification', 'User ID Validation failed', '<p>The user [UserName] failed the ID Validation process.</p><p>You can view the results of the ID validation here: \n\n [VerificationLink]</p>', 'enabled', 0, '2019-11-01 07:22:36', NULL, 924),
            (21, 'Password recovery', 'user', 'Password recovery', '[SiteName ] password recovery', '<p>[FirstName] [LastName],</p><p>Your confirmation code is [ConfirmationCode]</p>', 'enabled', 1, '2019-11-01 07:22:36', NULL, 916),
            (22, 'Incoming messages', 'admin', 'Message alerts', 'New private message at [SiteName].', 'Hi [PrivateMessageRecipient],\n\nThis is an automatic reminder from the site [SiteName]. You have received a new private message from [PrivateMessageAuthor].\n\n', 'enabled', 1, '2019-11-01 07:22:36', NULL, 912),
            (23, 'New transfer request (Admin)', 'admin', '', 'New transfer request created by admin', 'Admin user, [FirstName] [LastName], has requested a transfer for user, [UserName]. Please log in to your Admin interface to view the request.', 'enabled', 1, '2019-11-01 07:22:49', NULL, 947),
            (24, 'Dormant profile', 'admin', '', 'Dormant profile', '[UserName] profile status is now Dormant due to inactivity.\n[Logo]', 'enabled', 1, '2019-11-01 07:22:43', NULL, 944);

            CREATE TABLE `user_settings` (
              `id` int(11) UNSIGNED NOT NULL,
              `notification_name` varchar(255) NOT NULL,
              `uid` varchar(255) NOT NULL,
              `is_active` tinyint(1) NOT NULL DEFAULT '1'
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8;


            ALTER TABLE `schema_migrations`
              ADD PRIMARY KEY (`version`);

            ALTER TABLE `settings`
              ADD PRIMARY KEY (`id`),
              ADD UNIQUE KEY `uix_settings_name` (`name`);

            ALTER TABLE `templates`
              ADD PRIMARY KEY (`id`);

            ALTER TABLE `user_settings`
              ADD PRIMARY KEY (`id`),
              ADD UNIQUE KEY `uix_settings_notification_name_uid` (`notification_name`,`uid`);


            ALTER TABLE `settings`
              MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

            ALTER TABLE `templates`
              MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

            ALTER TABLE `user_settings`
              MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;
SQL;
    }
}
