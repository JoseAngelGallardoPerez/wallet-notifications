<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\DB;

class AddNotificationsSettingsForSmsProvider extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        DB::table('settings')->insert([
            'name' => 'africastalking_username',
            'value' => 'sandbox',
        ]);

        DB::table('settings')->insert([
            'name' => 'africastalking_api_code',
            'value' => '',
        ]);

        DB::table('settings')->insert([
            'name' => 'africastalking_short_code',
            'value' => '54231',
        ]);

        DB::table('settings')->insert([
            'name' => 'africastalking_env',
            'value' => 'sandbox',
        ]);

         DB::table('templates')->insert([
            'title' => 'Phone Verification',
            'scope' => 'user',
            'legend' => 'Phone Verification',
            'subject' => 'Phone Verification',
            'content' => '[ConfirmationCode]',
            'status' => 'enabled',
            'is_editable' => 1,
         ]);
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        //
    }
}
