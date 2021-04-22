<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\DB;

class AddSettingsPushStatus extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        $now = \Carbon\Carbon::now();
        DB::table('settings')->insert([
            'name' => 'push_status',
            'value' => 'disabled',
            'description' => 'Status of push notifications',
            'created_at' => $now,
            'updated_at' => $now,
        ]);
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        DB::delete("DELETE FROM settings WHERE `name` = ?", ["push_status"]);
    }
}
