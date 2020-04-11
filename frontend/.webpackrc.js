export default {
    'env': {
        'production': {
            'publicPath': '/static/',
            'html': {
                inject: true,
                template: './src/index.ejs',
            },
        }
    },
    'extraBabelPlugins': [
        [
            'import',
            {
                'libraryName': 'antd',
                'libraryDirectory': 'es',
                'style': 'css'
            }
        ]
    ]
}
