package gl

const (
	FALSE = 0
	TRUE  = 1

	/* ClearBufferMask */
	DEPTH_BUFFER_BIT   uint32 = 0x00000100
	STENCIL_BUFFER_BIT uint32 = 0x00000400
	COLOR_BUFFER_BIT   uint32 = 0x00004000

	/* BeginMode */
	POINTS         Enum = 0x0000
	LINES          Enum = 0x0001
	LINE_LOOP      Enum = 0x0002
	LINE_STRIP     Enum = 0x0003
	TRIANGLES      Enum = 0x0004
	TRIANGLE_STRIP Enum = 0x0005
	TRIANGLE_FAN   Enum = 0x0006

	/* AlphaFunction (not supported in ES20) */
	/*      NEVER */
	/*      LESS */
	/*      EQUAL */
	/*      LEQUAL */
	/*      GREATER */
	/*      NOTEQUAL */
	/*      GEQUAL */
	/*      ALWAYS */

	/* BlendingFactorDest */
	ZERO                Enum = 0
	ONE                 Enum = 1
	SRC_COLOR           Enum = 0x0300
	ONE_MINUS_SRC_COLOR Enum = 0x0301
	SRC_ALPHA           Enum = 0x0302
	ONE_MINUS_SRC_ALPHA Enum = 0x0303
	DST_ALPHA           Enum = 0x0304
	ONE_MINUS_DST_ALPHA Enum = 0x0305

	/* BlendingFactorSrc */
	/*      ZERO */
	/*      ONE */
	DST_COLOR           Enum = 0x0306
	ONE_MINUS_DST_COLOR Enum = 0x0307
	SRC_ALPHA_SATURATE  Enum = 0x0308
	/*      SRC_ALPHA */
	/*      ONE_MINUS_SRC_ALPHA */
	/*      DST_ALPHA */
	/*      ONE_MINUS_DST_ALPHA */

	/* BlendEquationSeparate */
	FUNC_ADD             Enum = 0x8006
	BLEND_EQUATION       Enum = 0x8009
	BLEND_EQUATION_RGB   Enum = 0x8009 /* same as BLEND_EQUATION */
	BLEND_EQUATION_ALPHA Enum = 0x883D

	/* BlendSubtract */
	FUNC_SUBTRACT         Enum = 0x800A
	FUNC_REVERSE_SUBTRACT Enum = 0x800B

	/* Separate Blend Functions */
	BLEND_DST_RGB            Enum = 0x80C8
	BLEND_SRC_RGB            Enum = 0x80C9
	BLEND_DST_ALPHA          Enum = 0x80CA
	BLEND_SRC_ALPHA          Enum = 0x80CB
	CONSTANT_COLOR           Enum = 0x8001
	ONE_MINUS_CONSTANT_COLOR Enum = 0x8002
	CONSTANT_ALPHA           Enum = 0x8003
	ONE_MINUS_CONSTANT_ALPHA Enum = 0x8004
	BLEND_COLOR              Enum = 0x8005

	/* Buffer Objects */
	ARRAY_BUFFER                 Enum = 0x8892
	ELEMENT_ARRAY_BUFFER         Enum = 0x8893
	ARRAY_BUFFER_BINDING         Enum = 0x8894
	ELEMENT_ARRAY_BUFFER_BINDING Enum = 0x8895

	STREAM_DRAW  Enum = 0x88E0
	STATIC_DRAW  Enum = 0x88E4
	DYNAMIC_DRAW Enum = 0x88E8

	BUFFER_SIZE  Enum = 0x8764
	BUFFER_USAGE Enum = 0x8765

	CURRENT_VERTEX_ATTRIB Enum = 0x8626

	/* CullFaceMode */
	FRONT          Enum = 0x0404
	BACK           Enum = 0x0405
	FRONT_AND_BACK Enum = 0x0408

	/* DepthFunction */
	/*      NEVER */
	/*      LESS */
	/*      EQUAL */
	/*      LEQUAL */
	/*      GREATER */
	/*      NOTEQUAL */
	/*      GEQUAL */
	/*      ALWAYS */

	/* EnableCap */
	/* TEXTURE_2D */
	CULL_FACE                Enum = 0x0B44
	BLEND                    Enum = 0x0BE2
	DITHER                   Enum = 0x0BD0
	STENCIL_TEST             Enum = 0x0B90
	DEPTH_TEST               Enum = 0x0B71
	SCISSOR_TEST             Enum = 0x0C11
	POLYGON_OFFSET_FILL      Enum = 0x8037
	SAMPLE_ALPHA_TO_COVERAGE Enum = 0x809E
	SAMPLE_COVERAGE          Enum = 0x80A0

	/* ErrorCode */
	NO_ERROR          Enum = 0
	INVALID_ENUM      Enum = 0x0500
	INVALID_VALUE     Enum = 0x0501
	INVALID_OPERATION Enum = 0x0502
	OUT_OF_MEMORY     Enum = 0x0505

	/* FrontFaceDirection */
	CW  Enum = 0x0900
	CCW Enum = 0x0901

	/* GetPName */
	LINE_WIDTH                   Enum = 0x0B21
	ALIASED_POINT_SIZE_RANGE     Enum = 0x846D
	ALIASED_LINE_WIDTH_RANGE     Enum = 0x846E
	CULL_FACE_MODE               Enum = 0x0B45
	FRONT_FACE                   Enum = 0x0B46
	DEPTH_RANGE                  Enum = 0x0B70
	DEPTH_WRITEMASK              Enum = 0x0B72
	DEPTH_CLEAR_VALUE            Enum = 0x0B73
	DEPTH_FUNC                   Enum = 0x0B74
	STENCIL_CLEAR_VALUE          Enum = 0x0B91
	STENCIL_FUNC                 Enum = 0x0B92
	STENCIL_FAIL                 Enum = 0x0B94
	STENCIL_PASS_DEPTH_FAIL      Enum = 0x0B95
	STENCIL_PASS_DEPTH_PASS      Enum = 0x0B96
	STENCIL_REF                  Enum = 0x0B97
	STENCIL_VALUE_MASK           Enum = 0x0B93
	STENCIL_WRITEMASK            Enum = 0x0B98
	STENCIL_BACK_FUNC            Enum = 0x8800
	STENCIL_BACK_FAIL            Enum = 0x8801
	STENCIL_BACK_PASS_DEPTH_FAIL Enum = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS Enum = 0x8803
	STENCIL_BACK_REF             Enum = 0x8CA3
	STENCIL_BACK_VALUE_MASK      Enum = 0x8CA4
	STENCIL_BACK_WRITEMASK       Enum = 0x8CA5
	VIEWPORT                     Enum = 0x0BA2
	SCISSOR_BOX                  Enum = 0x0C10
	/*      SCISSOR_TEST */
	COLOR_CLEAR_VALUE    Enum = 0x0C22
	COLOR_WRITEMASK      Enum = 0x0C23
	UNPACK_ALIGNMENT     Enum = 0x0CF5
	PACK_ALIGNMENT       Enum = 0x0D05
	MAX_TEXTURE_SIZE     Enum = 0x0D33
	MAX_VIEWPORT_DIMS    Enum = 0x0D3A
	SUBPIXEL_BITS        Enum = 0x0D50
	RED_BITS             Enum = 0x0D52
	GREEN_BITS           Enum = 0x0D53
	BLUE_BITS            Enum = 0x0D54
	ALPHA_BITS           Enum = 0x0D55
	DEPTH_BITS           Enum = 0x0D56
	STENCIL_BITS         Enum = 0x0D57
	POLYGON_OFFSET_UNITS Enum = 0x2A00
	/*      POLYGON_OFFSET_FILL */
	POLYGON_OFFSET_FACTOR  Enum = 0x8038
	TEXTURE_BINDING_2D     Enum = 0x8069
	SAMPLE_BUFFERS         Enum = 0x80A8
	SAMPLES                Enum = 0x80A9
	SAMPLE_COVERAGE_VALUE  Enum = 0x80AA
	SAMPLE_COVERAGE_INVERT Enum = 0x80AB

	/* GetTextureParameter */
	/*      TEXTURE_MAG_FILTER */
	/*      TEXTURE_MIN_FILTER */
	/*      TEXTURE_WRAP_S */
	/*      TEXTURE_WRAP_T */

	COMPRESSED_TEXTURE_FORMATS Enum = 0x86A3
	/* HintMode */
	DONT_CARE Enum = 0x1100
	FASTEST   Enum = 0x1101
	NICEST    Enum = 0x1102
	/* HintTarget */
	GENERATE_MIPMAP_HINT Enum = 0x8192
	/* DataType */
	BYTE           Enum = 0x1400
	UNSIGNED_BYTE  Enum = 0x1401
	SHORT          Enum = 0x1402
	UNSIGNED_SHORT Enum = 0x1403
	INT            Enum = 0x1404
	UNSIGNED_INT   Enum = 0x1405
	FLOAT          Enum = 0x1406
	/* PixelFormat */
	DEPTH_COMPONENT Enum = 0x1902
	ALPHA           Enum = 0x1906
	RGB             Enum = 0x1907
	RGBA            Enum = 0x1908
	LUMINANCE       Enum = 0x1909
	LUMINANCE_ALPHA Enum = 0x190A
	/* PixelType */
	/*      UNSIGNED_BYTE */
	UNSIGNED_SHORT_4_4_4_4 Enum = 0x8033
	UNSIGNED_SHORT_5_5_5_1 Enum = 0x8034
	UNSIGNED_SHORT_5_6_5   Enum = 0x8363
	/* Shaders */
	FRAGMENT_SHADER                  Enum = 0x8B30
	VERTEX_SHADER                    Enum = 0x8B31
	MAX_VERTEX_ATTRIBS               Enum = 0x8869
	MAX_VERTEX_UNIFORM_VECTORS       Enum = 0x8DFB
	MAX_VARYING_VECTORS              Enum = 0x8DFC
	MAX_COMBINED_TEXTURE_IMAGE_UNITS Enum = 0x8B4D
	MAX_VERTEX_TEXTURE_IMAGE_UNITS   Enum = 0x8B4C
	MAX_TEXTURE_IMAGE_UNITS          Enum = 0x8872
	MAX_FRAGMENT_UNIFORM_VECTORS     Enum = 0x8DFD
	SHADER_TYPE                      Enum = 0x8B4F
	DELETE_STATUS                    Enum = 0x8B80
	LINK_STATUS                      Enum = 0x8B82
	VALIDATE_STATUS                  Enum = 0x8B83
	ATTACHED_SHADERS                 Enum = 0x8B85
	ACTIVE_UNIFORMS                  Enum = 0x8B86
	ACTIVE_ATTRIBUTES                Enum = 0x8B89
	SHADING_LANGUAGE_VERSION         Enum = 0x8B8C
	CURRENT_PROGRAM                  Enum = 0x8B8D
	/* StencilFunction */
	NEVER    Enum = 0x0200
	LESS     Enum = 0x0201
	EQUAL    Enum = 0x0202
	LEQUAL   Enum = 0x0203
	GREATER  Enum = 0x0204
	NOTEQUAL Enum = 0x0205
	GEQUAL   Enum = 0x0206
	ALWAYS   Enum = 0x0207
	/* StencilOp */
	/*      ZERO */
	KEEP      Enum = 0x1E00
	REPLACE   Enum = 0x1E01
	INCR      Enum = 0x1E02
	DECR      Enum = 0x1E03
	INVERT    Enum = 0x150A
	INCR_WRAP Enum = 0x8507
	DECR_WRAP Enum = 0x8508
	/* StringName */
	VENDOR   Enum = 0x1F00
	RENDERER Enum = 0x1F01
	VERSION  Enum = 0x1F02
	/* TextureMagFilter */
	NEAREST Enum = 0x2600
	LINEAR  Enum = 0x2601
	/* TextureMinFilter */
	/*      NEAREST */
	/*      LINEAR */
	NEAREST_MIPMAP_NEAREST Enum = 0x2700
	LINEAR_MIPMAP_NEAREST  Enum = 0x2701
	NEAREST_MIPMAP_LINEAR  Enum = 0x2702
	LINEAR_MIPMAP_LINEAR   Enum = 0x2703
	/* TextureParameterName */
	TEXTURE_MAG_FILTER Enum = 0x2800
	TEXTURE_MIN_FILTER Enum = 0x2801
	TEXTURE_WRAP_S     Enum = 0x2802
	TEXTURE_WRAP_T     Enum = 0x2803
	/* TextureTarget */
	TEXTURE_2D                  Enum = 0x0DE1
	TEXTURE                     Enum = 0x1702
	TEXTURE_CUBE_MAP            Enum = 0x8513
	TEXTURE_BINDING_CUBE_MAP    Enum = 0x8514
	TEXTURE_CUBE_MAP_POSITIVE_X Enum = 0x8515
	TEXTURE_CUBE_MAP_NEGATIVE_X Enum = 0x8516
	TEXTURE_CUBE_MAP_POSITIVE_Y Enum = 0x8517
	TEXTURE_CUBE_MAP_NEGATIVE_Y Enum = 0x8518
	TEXTURE_CUBE_MAP_POSITIVE_Z Enum = 0x8519
	TEXTURE_CUBE_MAP_NEGATIVE_Z Enum = 0x851A
	MAX_CUBE_MAP_TEXTURE_SIZE   Enum = 0x851C
	/* TextureUnit */
	TEXTURE0       Enum = 0x84C0
	TEXTURE1       Enum = 0x84C1
	TEXTURE2       Enum = 0x84C2
	TEXTURE3       Enum = 0x84C3
	TEXTURE4       Enum = 0x84C4
	TEXTURE5       Enum = 0x84C5
	TEXTURE6       Enum = 0x84C6
	TEXTURE7       Enum = 0x84C7
	TEXTURE8       Enum = 0x84C8
	TEXTURE9       Enum = 0x84C9
	TEXTURE10      Enum = 0x84CA
	TEXTURE11      Enum = 0x84CB
	TEXTURE12      Enum = 0x84CC
	TEXTURE13      Enum = 0x84CD
	TEXTURE14      Enum = 0x84CE
	TEXTURE15      Enum = 0x84CF
	TEXTURE16      Enum = 0x84D0
	TEXTURE17      Enum = 0x84D1
	TEXTURE18      Enum = 0x84D2
	TEXTURE19      Enum = 0x84D3
	TEXTURE20      Enum = 0x84D4
	TEXTURE21      Enum = 0x84D5
	TEXTURE22      Enum = 0x84D6
	TEXTURE23      Enum = 0x84D7
	TEXTURE24      Enum = 0x84D8
	TEXTURE25      Enum = 0x84D9
	TEXTURE26      Enum = 0x84DA
	TEXTURE27      Enum = 0x84DB
	TEXTURE28      Enum = 0x84DC
	TEXTURE29      Enum = 0x84DD
	TEXTURE30      Enum = 0x84DE
	TEXTURE31      Enum = 0x84DF
	ACTIVE_TEXTURE Enum = 0x84E0
	/* TextureWrapMode */
	REPEAT          Enum = 0x2901
	CLAMP_TO_EDGE   Enum = 0x812F
	MIRRORED_REPEAT Enum = 0x8370
	/* Uniform Types */
	FLOAT_VEC2   Enum = 0x8B50
	FLOAT_VEC3   Enum = 0x8B51
	FLOAT_VEC4   Enum = 0x8B52
	INT_VEC2     Enum = 0x8B53
	INT_VEC3     Enum = 0x8B54
	INT_VEC4     Enum = 0x8B55
	BOOL         Enum = 0x8B56
	BOOL_VEC2    Enum = 0x8B57
	BOOL_VEC3    Enum = 0x8B58
	BOOL_VEC4    Enum = 0x8B59
	FLOAT_MAT2   Enum = 0x8B5A
	FLOAT_MAT3   Enum = 0x8B5B
	FLOAT_MAT4   Enum = 0x8B5C
	SAMPLER_2D   Enum = 0x8B5E
	SAMPLER_CUBE Enum = 0x8B60
	/* Vertex Arrays */
	VERTEX_ATTRIB_ARRAY_ENABLED        Enum = 0x8622
	VERTEX_ATTRIB_ARRAY_SIZE           Enum = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE         Enum = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE           Enum = 0x8625
	VERTEX_ATTRIB_ARRAY_NORMALIZED     Enum = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER        Enum = 0x8645
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING Enum = 0x889F
	/* Read Format */
	IMPLEMENTATION_COLOR_READ_TYPE   Enum = 0x8B9A
	IMPLEMENTATION_COLOR_READ_FORMAT Enum = 0x8B9B
	/* Shader Source */
	COMPILE_STATUS Enum = 0x8B81
	/* Shader Precision-Specified Types */
	LOW_FLOAT    Enum = 0x8DF0
	MEDIUM_FLOAT Enum = 0x8DF1
	HIGH_FLOAT   Enum = 0x8DF2
	LOW_INT      Enum = 0x8DF3
	MEDIUM_INT   Enum = 0x8DF4
	HIGH_INT     Enum = 0x8DF5
	/* Framebuffer Object. */
	FRAMEBUFFER                                   Enum = 0x8D40
	RENDERBUFFER                                  Enum = 0x8D41
	RGBA4                                         Enum = 0x8056
	RGB5_A1                                       Enum = 0x8057
	RGBA8                                         Enum = 0x8058
	RGB565                                        Enum = 0x8D62
	DEPTH_COMPONENT16                             Enum = 0x81A5
	STENCIL_INDEX8                                Enum = 0x8D48
	DEPTH_STENCIL                                 Enum = 0x84F9
	RENDERBUFFER_WIDTH                            Enum = 0x8D42
	RENDERBUFFER_HEIGHT                           Enum = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                  Enum = 0x8D44
	RENDERBUFFER_RED_SIZE                         Enum = 0x8D50
	RENDERBUFFER_GREEN_SIZE                       Enum = 0x8D51
	RENDERBUFFER_BLUE_SIZE                        Enum = 0x8D52
	RENDERBUFFER_ALPHA_SIZE                       Enum = 0x8D53
	RENDERBUFFER_DEPTH_SIZE                       Enum = 0x8D54
	RENDERBUFFER_STENCIL_SIZE                     Enum = 0x8D55
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE            Enum = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME            Enum = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL          Enum = 0x8CD2
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE  Enum = 0x8CD3
	COLOR_ATTACHMENT0                             Enum = 0x8CE0
	DEPTH_ATTACHMENT                              Enum = 0x8D00
	STENCIL_ATTACHMENT                            Enum = 0x8D20
	DEPTH_STENCIL_ATTACHMENT                      Enum = 0x821A
	NONE                                          Enum = 0
	FRAMEBUFFER_COMPLETE                          Enum = 0x8CD5
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT             Enum = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT     Enum = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS             Enum = 0x8CD9
	FRAMEBUFFER_UNSUPPORTED                       Enum = 0x8CDD
	FRAMEBUFFER_BINDING                           Enum = 0x8CA6
	RENDERBUFFER_BINDING                          Enum = 0x8CA7
	MAX_RENDERBUFFER_SIZE                         Enum = 0x84E8
	INVALID_FRAMEBUFFER_OPERATION                 Enum = 0x0506
	READ_BUFFER                                   Enum = 0x0C02
	UNPACK_ROW_LENGTH                             Enum = 0x0CF2
	UNPACK_SKIP_ROWS                              Enum = 0x0CF3
	UNPACK_SKIP_PIXELS                            Enum = 0x0CF4
	PACK_ROW_LENGTH                               Enum = 0x0D02
	PACK_SKIP_ROWS                                Enum = 0x0D03
	PACK_SKIP_PIXELS                              Enum = 0x0D04
	COLOR                                         Enum = 0x1800
	DEPTH                                         Enum = 0x1801
	STENCIL                                       Enum = 0x1802
	RED                                           Enum = 0x1903
	RGB8                                          Enum = 0x8051
	RGB10_A2                                      Enum = 0x8059
	TEXTURE_BINDING_3D                            Enum = 0x806A
	UNPACK_SKIP_IMAGES                            Enum = 0x806D
	UNPACK_IMAGE_HEIGHT                           Enum = 0x806E
	TEXTURE_3D                                    Enum = 0x806F
	TEXTURE_WRAP_R                                Enum = 0x8072
	MAX_3D_TEXTURE_SIZE                           Enum = 0x8073
	UNSIGNED_INT_2_10_10_10_REV                   Enum = 0x8368
	MAX_ELEMENTS_VERTICES                         Enum = 0x80E8
	MAX_ELEMENTS_INDICES                          Enum = 0x80E9
	TEXTURE_MIN_LOD                               Enum = 0x813A
	TEXTURE_MAX_LOD                               Enum = 0x813B
	TEXTURE_BASE_LEVEL                            Enum = 0x813C
	TEXTURE_MAX_LEVEL                             Enum = 0x813D
	MIN                                           Enum = 0x8007
	MAX                                           Enum = 0x8008
	DEPTH_COMPONENT24                             Enum = 0x81A6
	MAX_TEXTURE_LOD_BIAS                          Enum = 0x84FD
	TEXTURE_COMPARE_MODE                          Enum = 0x884C
	TEXTURE_COMPARE_FUNC                          Enum = 0x884D
	CURRENT_QUERY                                 Enum = 0x8865
	QUERY_RESULT                                  Enum = 0x8866
	QUERY_RESULT_AVAILABLE                        Enum = 0x8867
	STREAM_READ                                   Enum = 0x88E1
	STREAM_COPY                                   Enum = 0x88E2
	STATIC_READ                                   Enum = 0x88E5
	STATIC_COPY                                   Enum = 0x88E6
	DYNAMIC_READ                                  Enum = 0x88E9
	DYNAMIC_COPY                                  Enum = 0x88EA
	MAX_DRAW_BUFFERS                              Enum = 0x8824
	DRAW_BUFFER0                                  Enum = 0x8825
	DRAW_BUFFER1                                  Enum = 0x8826
	DRAW_BUFFER2                                  Enum = 0x8827
	DRAW_BUFFER3                                  Enum = 0x8828
	DRAW_BUFFER4                                  Enum = 0x8829
	DRAW_BUFFER5                                  Enum = 0x882A
	DRAW_BUFFER6                                  Enum = 0x882B
	DRAW_BUFFER7                                  Enum = 0x882C
	DRAW_BUFFER8                                  Enum = 0x882D
	DRAW_BUFFER9                                  Enum = 0x882E
	DRAW_BUFFER10                                 Enum = 0x882F
	DRAW_BUFFER11                                 Enum = 0x8830
	DRAW_BUFFER12                                 Enum = 0x8831
	DRAW_BUFFER13                                 Enum = 0x8832
	DRAW_BUFFER14                                 Enum = 0x8833
	DRAW_BUFFER15                                 Enum = 0x8834
	MAX_FRAGMENT_UNIFORM_COMPONENTS               Enum = 0x8B49
	MAX_VERTEX_UNIFORM_COMPONENTS                 Enum = 0x8B4A
	SAMPLER_3D                                    Enum = 0x8B5F
	SAMPLER_2D_SHADOW                             Enum = 0x8B62
	FRAGMENT_SHADER_DERIVATIVE_HINT               Enum = 0x8B8B
	PIXEL_PACK_BUFFER                             Enum = 0x88EB
	PIXEL_UNPACK_BUFFER                           Enum = 0x88EC
	PIXEL_PACK_BUFFER_BINDING                     Enum = 0x88ED
	PIXEL_UNPACK_BUFFER_BINDING                   Enum = 0x88EF
	FLOAT_MAT2x3                                  Enum = 0x8B65
	FLOAT_MAT2x4                                  Enum = 0x8B66
	FLOAT_MAT3x2                                  Enum = 0x8B67
	FLOAT_MAT3x4                                  Enum = 0x8B68
	FLOAT_MAT4x2                                  Enum = 0x8B69
	FLOAT_MAT4x3                                  Enum = 0x8B6A
	SRGB                                          Enum = 0x8C40
	SRGB8                                         Enum = 0x8C41
	SRGB8_ALPHA8                                  Enum = 0x8C43
	COMPARE_REF_TO_TEXTURE                        Enum = 0x884E
	RGBA32F                                       Enum = 0x8814
	RGB32F                                        Enum = 0x8815
	RGBA16F                                       Enum = 0x881A
	RGB16F                                        Enum = 0x881B
	VERTEX_ATTRIB_ARRAY_INTEGER                   Enum = 0x88FD
	MAX_ARRAY_TEXTURE_LAYERS                      Enum = 0x88FF
	MIN_PROGRAM_TEXEL_OFFSET                      Enum = 0x8904
	MAX_PROGRAM_TEXEL_OFFSET                      Enum = 0x8905
	MAX_VARYING_COMPONENTS                        Enum = 0x8B4B
	TEXTURE_2D_ARRAY                              Enum = 0x8C1A
	TEXTURE_BINDING_2D_ARRAY                      Enum = 0x8C1D
	R11F_G11F_B10F                                Enum = 0x8C3A
	UNSIGNED_INT_10F_11F_11F_REV                  Enum = 0x8C3B
	RGB9_E5                                       Enum = 0x8C3D
	UNSIGNED_INT_5_9_9_9_REV                      Enum = 0x8C3E
	TRANSFORM_FEEDBACK_BUFFER_MODE                Enum = 0x8C7F
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS    Enum = 0x8C80
	TRANSFORM_FEEDBACK_VARYINGS                   Enum = 0x8C83
	TRANSFORM_FEEDBACK_BUFFER_START               Enum = 0x8C84
	TRANSFORM_FEEDBACK_BUFFER_SIZE                Enum = 0x8C85
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN         Enum = 0x8C88
	RASTERIZER_DISCARD                            Enum = 0x8C89
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS Enum = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS       Enum = 0x8C8B
	INTERLEAVED_ATTRIBS                           Enum = 0x8C8C
	SEPARATE_ATTRIBS                              Enum = 0x8C8D
	TRANSFORM_FEEDBACK_BUFFER                     Enum = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_BINDING             Enum = 0x8C8F
	RGBA32UI                                      Enum = 0x8D70
	RGB32UI                                       Enum = 0x8D71
	RGBA16UI                                      Enum = 0x8D76
	RGB16UI                                       Enum = 0x8D77
	RGBA8UI                                       Enum = 0x8D7C
	RGB8UI                                        Enum = 0x8D7D
	RGBA32I                                       Enum = 0x8D82
	RGB32I                                        Enum = 0x8D83
	RGBA16I                                       Enum = 0x8D88
	RGB16I                                        Enum = 0x8D89
	RGBA8I                                        Enum = 0x8D8E
	RGB8I                                         Enum = 0x8D8F
	RED_INTEGER                                   Enum = 0x8D94
	RGB_INTEGER                                   Enum = 0x8D98
	RGBA_INTEGER                                  Enum = 0x8D99
	SAMPLER_2D_ARRAY                              Enum = 0x8DC1
	SAMPLER_2D_ARRAY_SHADOW                       Enum = 0x8DC4
	SAMPLER_CUBE_SHADOW                           Enum = 0x8DC5
	UNSIGNED_INT_VEC2                             Enum = 0x8DC6
	UNSIGNED_INT_VEC3                             Enum = 0x8DC7
	UNSIGNED_INT_VEC4                             Enum = 0x8DC8
	INT_SAMPLER_2D                                Enum = 0x8DCA
	INT_SAMPLER_3D                                Enum = 0x8DCB
	INT_SAMPLER_CUBE                              Enum = 0x8DCC
	INT_SAMPLER_2D_ARRAY                          Enum = 0x8DCF
	UNSIGNED_INT_SAMPLER_2D                       Enum = 0x8DD2
	UNSIGNED_INT_SAMPLER_3D                       Enum = 0x8DD3
	UNSIGNED_INT_SAMPLER_CUBE                     Enum = 0x8DD4
	UNSIGNED_INT_SAMPLER_2D_ARRAY                 Enum = 0x8DD7
	DEPTH_COMPONENT32F                            Enum = 0x8CAC
	DEPTH32F_STENCIL8                             Enum = 0x8CAD
	FLOAT_32_UNSIGNED_INT_24_8_REV                Enum = 0x8DAD
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING         Enum = 0x8210
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE         Enum = 0x8211
	FRAMEBUFFER_ATTACHMENT_RED_SIZE               Enum = 0x8212
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE             Enum = 0x8213
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE              Enum = 0x8214
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE             Enum = 0x8215
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE             Enum = 0x8216
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE           Enum = 0x8217
	FRAMEBUFFER_DEFAULT                           Enum = 0x8218
	UNSIGNED_INT_24_8                             Enum = 0x84FA
	DEPTH24_STENCIL8                              Enum = 0x88F0
	UNSIGNED_NORMALIZED                           Enum = 0x8C17
	DRAW_FRAMEBUFFER_BINDING                      Enum = 0x8CA6 /* Same as FRAMEBUFFER_BINDING */
	READ_FRAMEBUFFER                              Enum = 0x8CA8
	DRAW_FRAMEBUFFER                              Enum = 0x8CA9
	READ_FRAMEBUFFER_BINDING                      Enum = 0x8CAA
	RENDERBUFFER_SAMPLES                          Enum = 0x8CAB
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER          Enum = 0x8CD4
	MAX_COLOR_ATTACHMENTS                         Enum = 0x8CDF
	COLOR_ATTACHMENT1                             Enum = 0x8CE1
	COLOR_ATTACHMENT2                             Enum = 0x8CE2
	COLOR_ATTACHMENT3                             Enum = 0x8CE3
	COLOR_ATTACHMENT4                             Enum = 0x8CE4
	COLOR_ATTACHMENT5                             Enum = 0x8CE5
	COLOR_ATTACHMENT6                             Enum = 0x8CE6
	COLOR_ATTACHMENT7                             Enum = 0x8CE7
	COLOR_ATTACHMENT8                             Enum = 0x8CE8
	COLOR_ATTACHMENT9                             Enum = 0x8CE9
	COLOR_ATTACHMENT10                            Enum = 0x8CEA
	COLOR_ATTACHMENT11                            Enum = 0x8CEB
	COLOR_ATTACHMENT12                            Enum = 0x8CEC
	COLOR_ATTACHMENT13                            Enum = 0x8CED
	COLOR_ATTACHMENT14                            Enum = 0x8CEE
	COLOR_ATTACHMENT15                            Enum = 0x8CEF
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE            Enum = 0x8D56
	MAX_SAMPLES                                   Enum = 0x8D57
	HALF_FLOAT                                    Enum = 0x140B
	RG                                            Enum = 0x8227
	RG_INTEGER                                    Enum = 0x8228
	R8                                            Enum = 0x8229
	RG8                                           Enum = 0x822B
	R16F                                          Enum = 0x822D
	R32F                                          Enum = 0x822E
	RG16F                                         Enum = 0x822F
	RG32F                                         Enum = 0x8230
	R8I                                           Enum = 0x8231
	R8UI                                          Enum = 0x8232
	R16I                                          Enum = 0x8233
	R16UI                                         Enum = 0x8234
	R32I                                          Enum = 0x8235
	R32UI                                         Enum = 0x8236
	RG8I                                          Enum = 0x8237
	RG8UI                                         Enum = 0x8238
	RG16I                                         Enum = 0x8239
	RG16UI                                        Enum = 0x823A
	RG32I                                         Enum = 0x823B
	RG32UI                                        Enum = 0x823C
	VERTEX_ARRAY_BINDING                          Enum = 0x85B5
	R8_SNORM                                      Enum = 0x8F94
	RG8_SNORM                                     Enum = 0x8F95
	RGB8_SNORM                                    Enum = 0x8F96
	RGBA8_SNORM                                   Enum = 0x8F97
	SIGNED_NORMALIZED                             Enum = 0x8F9C
	COPY_READ_BUFFER                              Enum = 0x8F36
	COPY_WRITE_BUFFER                             Enum = 0x8F37
	COPY_READ_BUFFER_BINDING                      Enum = 0x8F36 /* Same as COPY_READ_BUFFER */
	COPY_WRITE_BUFFER_BINDING                     Enum = 0x8F37 /* Same as COPY_WRITE_BUFFER */
	UNIFORM_BUFFER                                Enum = 0x8A11
	UNIFORM_BUFFER_BINDING                        Enum = 0x8A28
	UNIFORM_BUFFER_START                          Enum = 0x8A29
	UNIFORM_BUFFER_SIZE                           Enum = 0x8A2A
	MAX_VERTEX_UNIFORM_BLOCKS                     Enum = 0x8A2B
	MAX_FRAGMENT_UNIFORM_BLOCKS                   Enum = 0x8A2D
	MAX_COMBINED_UNIFORM_BLOCKS                   Enum = 0x8A2E
	MAX_UNIFORM_BUFFER_BINDINGS                   Enum = 0x8A2F
	MAX_UNIFORM_BLOCK_SIZE                        Enum = 0x8A30
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS        Enum = 0x8A31
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS      Enum = 0x8A33
	UNIFORM_BUFFER_OFFSET_ALIGNMENT               Enum = 0x8A34
	ACTIVE_UNIFORM_BLOCKS                         Enum = 0x8A36
	UNIFORM_TYPE                                  Enum = 0x8A37
	UNIFORM_SIZE                                  Enum = 0x8A38
	UNIFORM_BLOCK_INDEX                           Enum = 0x8A3A
	UNIFORM_OFFSET                                Enum = 0x8A3B
	UNIFORM_ARRAY_STRIDE                          Enum = 0x8A3C
	UNIFORM_MATRIX_STRIDE                         Enum = 0x8A3D
	UNIFORM_IS_ROW_MAJOR                          Enum = 0x8A3E
	UNIFORM_BLOCK_BINDING                         Enum = 0x8A3F
	UNIFORM_BLOCK_DATA_SIZE                       Enum = 0x8A40
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                 Enum = 0x8A42
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES          Enum = 0x8A43
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER     Enum = 0x8A44
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER   Enum = 0x8A46
	INVALID_INDEX                                 Enum = 0xFFFFFFFF
	MAX_VERTEX_OUTPUT_COMPONENTS                  Enum = 0x9122
	MAX_FRAGMENT_INPUT_COMPONENTS                 Enum = 0x9125
	MAX_SERVER_WAIT_TIMEOUT                       Enum = 0x9111
	OBJECT_TYPE                                   Enum = 0x9112
	SYNC_CONDITION                                Enum = 0x9113
	SYNC_STATUS                                   Enum = 0x9114
	SYNC_FLAGS                                    Enum = 0x9115
	SYNC_FENCE                                    Enum = 0x9116
	SYNC_GPU_COMMANDS_COMPLETE                    Enum = 0x9117
	UNSIGNALED                                    Enum = 0x9118
	SIGNALED                                      Enum = 0x9119
	ALREADY_SIGNALED                              Enum = 0x911A
	TIMEOUT_EXPIRED                               Enum = 0x911B
	CONDITION_SATISFIED                           Enum = 0x911C
	WAIT_FAILED                                   Enum = 0x911D
	SYNC_FLUSH_COMMANDS_BIT                       Enum = 0x00000001
	VERTEX_ATTRIB_ARRAY_DIVISOR                   Enum = 0x88FE
	ANY_SAMPLES_PASSED                            Enum = 0x8C2F
	ANY_SAMPLES_PASSED_CONSERVATIVE               Enum = 0x8D6A
	SAMPLER_BINDING                               Enum = 0x8919
	RGB10_A2UI                                    Enum = 0x906F
	INT_2_10_10_10_REV                            Enum = 0x8D9F
	TRANSFORM_FEEDBACK                            Enum = 0x8E22
	TRANSFORM_FEEDBACK_PAUSED                     Enum = 0x8E23
	TRANSFORM_FEEDBACK_ACTIVE                     Enum = 0x8E24
	TRANSFORM_FEEDBACK_BINDING                    Enum = 0x8E25
	TEXTURE_IMMUTABLE_FORMAT                      Enum = 0x912F
	MAX_ELEMENT_INDEX                             Enum = 0x8D6B
	TEXTURE_IMMUTABLE_LEVELS                      Enum = 0x82DF
)
