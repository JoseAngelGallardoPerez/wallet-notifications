<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Doctrine\DBAL\Types\StringType;
use Doctrine\DBAL\Types\Type;

class CreatePushTokens extends Migration
{
    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function up()
    {
        if (!Type::hasType('enum')) {
            Type::addType('enum', StringType::class);
        }
        DB::connection()->getDoctrineSchemaManager()->getDatabasePlatform()->registerDoctrineTypeMapping('enum', 'string');

        Schema::create('push_tokens', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->string("push_token")->nullable(false)->primary();
            $table->enum('os', ['ios', 'android', 'other'])->default("other");
            $table->string("name")->default("");
            $table->string("uid")->nullable(false)->index();
            $table->string("device_id")->nullable(false)->default("");

            $table->dateTime('created_at')->nullable(false);
            $table->dateTime('updated_at')->nullable(false);
            //$table->timestamps();
        });
    }

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('push_tokens');
    }
}
