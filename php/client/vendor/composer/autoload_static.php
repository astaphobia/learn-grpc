<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInitb4756bbec58c47fb707b0ab314b962d0
{
    public static $files = array (
        '7c623dc93c03ddd466695ebe6517b393' => __DIR__ . '/..' . '/google/protobuf/php/src/Google/Protobuf/descriptor.php',
    );

    public static $prefixLengthsPsr4 = array (
        'G' => 
        array (
            'Grpc\\' => 5,
            'Google\\Protobuf\\Internal\\' => 25,
            'GPBMetadata\\Google\\Protobuf\\Internal\\' => 37,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'Grpc\\' => 
        array (
            0 => __DIR__ . '/..' . '/grpc/grpc/src/php/lib/Grpc',
        ),
        'Google\\Protobuf\\Internal\\' => 
        array (
            0 => __DIR__ . '/..' . '/google/protobuf/php/src/Google/Protobuf/Internal',
        ),
        'GPBMetadata\\Google\\Protobuf\\Internal\\' => 
        array (
            0 => __DIR__ . '/..' . '/google/protobuf/php/src/GPBMetadata/Google/Protobuf/Internal',
        ),
    );

    public static $fallbackDirsPsr4 = array (
        0 => __DIR__ . '/../..' . '/../protobuf',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInitb4756bbec58c47fb707b0ab314b962d0::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInitb4756bbec58c47fb707b0ab314b962d0::$prefixDirsPsr4;
            $loader->fallbackDirsPsr4 = ComposerStaticInitb4756bbec58c47fb707b0ab314b962d0::$fallbackDirsPsr4;

        }, null, ClassLoader::class);
    }
}