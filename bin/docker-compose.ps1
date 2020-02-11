# docker-compose のショートカット
# => 一部コマンドのみ内容を切り替えたりする用

$rootDir = $PSScriptRoot | Split-Path
$dockerDir = Join-Path $rootDir docker

function Build-Args($params) {
    $dev = 'dev'
    if (!$params) {
        return @()
    }
    if ($params.Length -eq 1) {
        switch ($params[0]) {
            'up' {
                return @('up', '-d', '--build')
            }
            'attach' {
                return @('exec', $dev, '/bin/bash')
            }
        }
    }
    return $params
}

try {
    Push-Location $dockerDir > $null
    $arguments = Build-Args $args
    docker-compose $arguments
}
finally {
    Pop-Location > $null
}