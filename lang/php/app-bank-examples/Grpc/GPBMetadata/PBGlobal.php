<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: global.proto

namespace GPBMetadata;

class PBGlobal
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\GPBEmpty::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0ace010a0c676c6f62616c2e70726f746f225a0a0b56657273696f6e496e" .
            "666f120f0a0776657273696f6e180120012809120f0a0773657276696365" .
            "18022001280912130a0b6465736372697074696f6e18032001280912140a" .
            "0c6163636573735f7269676874180420012809425a0a1e636f6d2e616b65" .
            "792e6f70656e2e626173652e677270632e676c6f62616c420b476c6f6261" .
            "6c50726f746f50015a22616b65792d6170702d62616e6b2d6578616d706c" .
            "65732f677270632f676c6f62616ca20204414b6579620670726f746f33"
        ));

        static::$is_initialized = true;
    }
}

