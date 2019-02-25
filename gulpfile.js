const { src, dest, parallel, series } = require('gulp');
const rollup = require('gulp-better-rollup');
const { uglify } = require('rollup-plugin-uglify');
const babel = require('rollup-plugin-babel');
const { ugh } = require('uglify-es').minify;
const flatten = require('gulp-flatten');
const glob = require('glob');
const path = require('path');

function js() {
    const entries = getJSEntries();

    return entries.map((name) => {
        const rollupOpts = {
            entry: 'static/js/' + name + '.entry.js',
            external: ['jquery'],
            output: {
                name: name,
                globals: {
                    jquery: 'jquery'
                },
                paths: {
                    jquery: 'https://code.jquery.com/jquery-3.2.1.min.js'
                },
                format: 'iife',
            },
            plugins: [
                babel({
                    exclude: 'node_modules/**'
                }),
                uglify({}, ugh)
            ]
        };

        return src('static/js/*.entry.js')
            .pipe(rollup(rollupOpts, 'iife'))
            .pipe(flatten())
            .pipe(dest('dist/js'));
    });
}

function getJSEntries() {
    let result = [];

    const entryPattrn = `static/js/*.entry.js`;
    const filenames = getFileNames(entryPattrn);

    filenames.forEach((name) => {
        result.push(name)
    })

    return result;
}

function getFileNames(pattern) {
    return glob.sync(pattern)
        .map((componentDir) => {
            console.log(componentDir);
            return path.basename(componentDir);
        });
}

exports.js = series(js)
exports.default = parallel(js)