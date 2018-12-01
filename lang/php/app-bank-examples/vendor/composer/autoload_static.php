<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInit1e2b409e96fba9b544771e42da91b205
{
    public static $prefixLengthsPsr4 = array (
        'G' => 
        array (
            'Grpc\\' => 5,
            'Google\\Protobuf\\' => 16,
            'GPBMetadata\\Google\\Protobuf\\' => 28,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'Grpc\\' => 
        array (
            0 => __DIR__ . '/..' . '/grpc/grpc/src/lib',
        ),
        'Google\\Protobuf\\' => 
        array (
            0 => __DIR__ . '/..' . '/google/protobuf/php/src/Google/Protobuf',
        ),
        'GPBMetadata\\Google\\Protobuf\\' => 
        array (
            0 => __DIR__ . '/..' . '/google/protobuf/php/src/GPBMetadata/Google/Protobuf',
        ),
    );

    public static $fallbackDirsPsr4 = array (
        0 => __DIR__ . '/../..' . '/Grpc',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInit1e2b409e96fba9b544771e42da91b205::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInit1e2b409e96fba9b544771e42da91b205::$prefixDirsPsr4;
            $loader->fallbackDirsPsr4 = ComposerStaticInit1e2b409e96fba9b544771e42da91b205::$fallbackDirsPsr4;

        }, null, ClassLoader::class);
    }
}