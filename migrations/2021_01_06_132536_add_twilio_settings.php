<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddTwilioSettings extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        $data = [
            [
                'name' => 'twilio_account_sid',
                'value' => '',
                'description' => 'Twilio account string identifier',
            ],
            [
                'name' => 'twilio_auth_token',
                'value' => '',
                'description' => 'Twilio account authentication token',
            ],
            [
                'name' => 'twilio_sms_from',
                'value' => '',
                'description' => 'Twilio registered phone number with "SMS" permissions',
            ],
            [
                'name' => 'sms_provider',
                'value' => 'twilio',
                'description' => 'Default SMS provider',
            ],
        ];

        DB::table('settings')->insert($data);
    }

    public function down()
    {
        $newParams = [
            'twilio_account_sid',
            'twilio_auth_token',
            'twilio_sms_from',
            'sms_provider',
        ];

        DB::table('settings')->whereIn('name', $newParams)->delete();
    }
}