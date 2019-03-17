const { src, dest, parallel, task } = require('gulp');
const rollup = require('gulp-better-rollup');
const { uglify } = require('rollup-plugin-uglify');
const babel = require('rollup-plugin-babel');
const { ugh } = require('uglify-es').minify;
const flatten = require('gulp-flatten');
const glob = require('glob');
const path = require('path');

function loadJSEntry(cb) {
    const entries = getJSEntries();

    return Promise.resolve(entries.map((name) => {
        return js(name);
    }));
}

function js(name) {
    const fullName = 'static/js/' + name;
    const rollupOpts = {
        entry: fullName,
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

    return src(fullName)
        .pipe(rollup(rollupOpts, 'iife'))
        .pipe(flatten())
        .pipe(dest('dist/js'));

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

exports.default = parallel(loadJSEntry);