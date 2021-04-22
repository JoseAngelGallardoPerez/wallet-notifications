<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class EmailVerificationTemplate extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        DB::table('templates')->insert([
            'title' => 'Email Confirmation',
            'scope' => 'user',
            'legend' => 'Email Confirmation',
            'subject' => 'Email Confirmation',
            'content' => 'Confirmation code: [ConfirmationCode]',
            'status' => 'enabled',
        ]);
        DB::table('templates')->where('title','Registration Confirmation')->update([
            'status' => 'disabled',
            'is_editable' => 0,
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
