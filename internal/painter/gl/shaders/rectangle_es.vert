#version 100

#ifdef GL_ES
# ifdef GL_FRAGMENT_PRECISION_HIGH
precision highp float;
# else
precision mediump float;
#endif
precision mediump int;
precision lowp sampler2D;
#endif

attribute vec2 vert;
attribute vec2 normal;
attribute float colorSwitch;
attribute float lineWidth;
attribute float feather;

varying vec2 delta;
varying float switch_var;
varying float lineWidth_var;
varying float feather_var;

void main() {
    switch_var = colorSwitch;
    lineWidth_var = lineWidth;
    feather_var = feather;
    if ( normal == vec2(0.0, 0.0) ) {
        gl_Position = vec4(vert, 0, 1);
    } else {
        delta = normal * lineWidth_var;
        gl_Position = vec4(vert + delta, 0, 1);
    }
}
