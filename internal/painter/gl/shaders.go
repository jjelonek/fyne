// auto-generated
// Code generated by '$ fyne bundle'. DO NOT EDIT.

package gl

import "fyne.io/fyne/v2"

var shaderLineFrag = &fyne.StaticResource{
	StaticName: "line.frag",
	StaticContent: []byte(
		"#version 110\n\nuniform vec4 color;\nuniform float lineWidth;\nuniform float feather;\n\nvarying vec2 delta;\n\nvoid main() {\n    float alpha = color.a;\n    float distance = length(delta);\n\n    if (feather == 0.0 || distance <= lineWidth - feather) {\n        gl_FragColor = color;\n    } else {\n        gl_FragColor = vec4(color.r, color.g, color.b, mix(color.a, 0.0, (distance - (lineWidth - feather)) / feather));\n    }\n}\n"),
}
var shaderLineVert = &fyne.StaticResource{
	StaticName: "line.vert",
	StaticContent: []byte(
		"#version 110\n\nattribute vec2 vert;\nattribute vec2 normal;\n    \nuniform float lineWidth;\n\nvarying vec2 delta;\n\nvoid main() {\n    delta = normal * lineWidth;\n\n    gl_Position = vec4(vert + delta, 0, 1);\n}\n"),
}
var shaderLineesFrag = &fyne.StaticResource{
	StaticName: "line_es.frag",
	StaticContent: []byte(
		"#version 100\n\n#ifdef GL_ES\n# ifdef GL_FRAGMENT_PRECISION_HIGH\nprecision highp float;\n# else\nprecision mediump float;\n#endif\nprecision mediump int;\nprecision lowp sampler2D;\n#endif\n\nuniform vec4 color;\nuniform float lineWidth;\nuniform float feather;\n\nvarying vec2 delta;\n\nvoid main() {\n    float alpha = color.a;\n    float distance = length(delta);\n\n    if (feather == 0.0 || distance <= lineWidth - feather) {\n        gl_FragColor = color;\n    } else {\n        gl_FragColor = vec4(color.r, color.g, color.b, mix(color.a, 0.0, (distance - (lineWidth - feather)) / feather));\n    }\n}\n"),
}
var shaderLineesVert = &fyne.StaticResource{
	StaticName: "line_es.vert",
	StaticContent: []byte(
		"#version 100\n\n#ifdef GL_ES\n# ifdef GL_FRAGMENT_PRECISION_HIGH\nprecision highp float;\n# else\nprecision mediump float;\n#endif\nprecision mediump int;\nprecision lowp sampler2D;\n#endif\n\nattribute vec2 vert;\nattribute vec2 normal;\n    \nuniform float lineWidth;\n\nvarying vec2 delta;\n\nvoid main() {\n    delta = normal * lineWidth;\n\n    gl_Position = vec4(vert + delta, 0, 1);\n}\n"),
}
var shaderSimpleFrag = &fyne.StaticResource{
	StaticName: "simple.frag",
	StaticContent: []byte(
		"#version 110\n\nuniform sampler2D tex;\n\nvarying vec2 fragTexCoord;\n\nvoid main() {\n    gl_FragColor = texture2D(tex, fragTexCoord);\n}\n"),
}
var shaderSimpleVert = &fyne.StaticResource{
	StaticName: "simple.vert",
	StaticContent: []byte(
		"#version 110\n\nattribute vec3 vert;\nattribute vec2 vertTexCoord;\nvarying vec2 fragTexCoord;\n\nvoid main() {\n    fragTexCoord = vertTexCoord;\n\n    gl_Position = vec4(vert, 1);\n}"),
}
var shaderSimpleesFrag = &fyne.StaticResource{
	StaticName: "simple_es.frag",
	StaticContent: []byte(
		"#version 100\n\n#ifdef GL_ES\n# ifdef GL_FRAGMENT_PRECISION_HIGH\nprecision highp float;\n# else\nprecision mediump float;\n#endif\nprecision mediump int;\nprecision lowp sampler2D;\n#endif\n\nuniform sampler2D tex;\n\nvarying vec2 fragTexCoord;\n\nvoid main() {\n    gl_FragColor = texture2D(tex, fragTexCoord);\n}\n"),
}
var shaderSimpleesVert = &fyne.StaticResource{
	StaticName: "simple_es.vert",
	StaticContent: []byte(
		"#version 100\n\n#ifdef GL_ES\n# ifdef GL_FRAGMENT_PRECISION_HIGH\nprecision highp float;\n# else\nprecision mediump float;\n#endif\nprecision mediump int;\nprecision lowp sampler2D;\n#endif\n\nattribute vec3 vert;\nattribute vec2 vertTexCoord;\nvarying vec2 fragTexCoord;\n\nvoid main() {\n    fragTexCoord = vertTexCoord;\n\n    gl_Position = vec4(vert, 1);\n}"),
}
var shaderRectangleFrag = &fyne.StaticResource{
	StaticName: "rectangle.frag",
	StaticContent: []byte(
		"#version 110\n\nuniform vec4 fill_color;\nuniform vec4 stroke_color;\n\nvarying vec2 delta;\nvarying float switch_var;\nvarying float lineWidth_var;\nvarying float feather_var;\n\nvoid main() {\n    vec4 color = vec4(1.0, 0.0, 0.0, 255.0);\n\n    if ( switch_var == 1.0 ){\n        color = fill_color;\n    }; \n    if ( switch_var == 2.0 ){\t\n        color = stroke_color;\n    };\n    if ( switch_var > 50.0 ){\t\n        color = vec4(stroke_color.r, stroke_color.g, stroke_color.b, stroke_color.a * switch_var / 100.0 );\n    };\n    if ( switch_var >= 10.0 && switch_var <= 50.0 ){\t\n        color = vec4(fill_color.r, fill_color.g, fill_color.b, fill_color.a * switch_var / 100.0 );\n    };\n    if ( delta != vec2(0.0, 0.0) ){\n        float distance = length(delta);\n\n        if (distance <= lineWidth_var - feather_var) {\n            gl_FragColor = color;\n        } else {\n            gl_FragColor = vec4(color.r, color.g, color.b, mix(color.a, 0.0, (distance - (lineWidth_var - feather_var)) / feather_var));\n        }\n    } else {\n        gl_FragColor = color;\n    }\n}\n"),
}
var shaderRectangleVert = &fyne.StaticResource{
	StaticName: "rectangle.vert",
	StaticContent: []byte(
		"#version 110\n\nattribute vec2 vert;\nattribute vec2 normal;\nattribute float colorSwitch;\nattribute float lineWidth;\nattribute float feather;\n\nvarying vec2 delta;\nvarying float switch_var;\nvarying float lineWidth_var;\nvarying float feather_var;\n\nvoid main() {\n    switch_var = colorSwitch;\n    lineWidth_var = lineWidth;\n    feather_var = feather;\n    if ( normal == vec2(0.0, 0.0) ) {\n        gl_Position = vec4(vert, 0, 1);\n    } else {\n        delta = normal * lineWidth_var;\n        gl_Position = vec4(vert + delta, 0, 1);\n    }\n}"),
}
var shaderRectangleesFrag = &fyne.StaticResource{
	StaticName: "rectangle_es.frag",
	StaticContent: []byte(
		"#version 100\n\n#ifdef GL_ES\n# ifdef GL_FRAGMENT_PRECISION_HIGH\nprecision highp float;\n# else\nprecision mediump float;\n#endif\nprecision mediump int;\nprecision lowp sampler2D;\n#endif\n\nuniform vec4 fill_color;\nuniform vec4 stroke_color;\n\nvarying vec2 delta;\nvarying float switch_var;\nvarying float lineWidth_var;\nvarying float feather_var;\nvec4 color;\n\nvoid main() {\n    if ( switch_var == 1.0 ){\n        color = fill_color;\n    }; \n    if ( switch_var == 2.0 ){\t\n        color = stroke_color;\n    };\n    if ( switch_var > 50.0 ){\t\n        color = vec4(stroke_color.r, stroke_color.g, stroke_color.b, stroke_color.a * switch_var / 100.0 );\n    };\n    if ( switch_var >= 10.0 && switch_var <= 50.0 ){\t\n        color = vec4(fill_color.r, fill_color.g, fill_color.b, fill_color.a * switch_var / 100.0 );\n    };\n    if ( delta != vec2(0.0, 0.0) ){\n        float distance = length(delta);\n\n        if (distance <= lineWidth_var - feather_var) {\n            gl_FragColor = color;\n        } else {\n            gl_FragColor = vec4(color.r, color.g, color.b, mix(color.a, 0.0, (distance - (lineWidth_var - feather_var)) / feather_var));\n        }\n    } else {\n        gl_FragColor = color;\n    }\n}"),
}
var shaderRectangleesVert = &fyne.StaticResource{
	StaticName: "rectangle_es.vert",
	StaticContent: []byte(
		"#version 100\n\n#ifdef GL_ES\n# ifdef GL_FRAGMENT_PRECISION_HIGH\nprecision highp float;\n# else\nprecision mediump float;\n#endif\nprecision mediump int;\nprecision lowp sampler2D;\n#endif\n\nattribute vec2 vert;\nattribute vec2 normal;\nattribute float colorSwitch;\nattribute float lineWidth;\nattribute float feather;\n\nvarying vec2 delta;\nvarying float switch_var;\nvarying float lineWidth_var;\nvarying float feather_var;\n\nvoid main() {\n    switch_var = colorSwitch;\n    lineWidth_var = lineWidth;\n    feather_var = feather;\n    if ( normal == vec2(0.0, 0.0) ) {\n        gl_Position = vec4(vert, 0, 1);\n    } else {\n        delta = normal * lineWidth_var;\n        gl_Position = vec4(vert + delta, 0, 1);\n    }\n}"),
}
