<?php


function reverst($data)
{
    for ($i = 1; $i < $data; $i++) {

        if ($i % 3 == 0 && $i % 5 == 0) {
            echo "tokopedia" . PHP_EOL;
        } elseif ($i % 3 == 0) {
            echo "toko" . PHP_EOL;
        } elseif ($i % 5 == 0) {
            echo "pedia" . PHP_EOL;
        } else {
            echo $i . PHP_EOL;
        }
    }
}

reverst(3);
