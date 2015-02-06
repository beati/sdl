package gameio

/*
#include <SDL2/SDL.h>
*/
import "C"

const (
	Key_UNKNOWN        = C.SDL_SCANCODE_UNKNOWN
	Key_A              = C.SDL_SCANCODE_A
	Key_B              = C.SDL_SCANCODE_B
	Key_C              = C.SDL_SCANCODE_C
	Key_D              = C.SDL_SCANCODE_D
	Key_E              = C.SDL_SCANCODE_E
	Key_F              = C.SDL_SCANCODE_F
	Key_G              = C.SDL_SCANCODE_G
	Key_H              = C.SDL_SCANCODE_H
	Key_I              = C.SDL_SCANCODE_I
	Key_J              = C.SDL_SCANCODE_J
	Key_K              = C.SDL_SCANCODE_K
	Key_L              = C.SDL_SCANCODE_L
	Key_M              = C.SDL_SCANCODE_M
	Key_N              = C.SDL_SCANCODE_N
	Key_O              = C.SDL_SCANCODE_O
	Key_P              = C.SDL_SCANCODE_P
	Key_Q              = C.SDL_SCANCODE_Q
	Key_R              = C.SDL_SCANCODE_R
	Key_S              = C.SDL_SCANCODE_S
	Key_T              = C.SDL_SCANCODE_T
	Key_U              = C.SDL_SCANCODE_U
	Key_V              = C.SDL_SCANCODE_V
	Key_W              = C.SDL_SCANCODE_W
	Key_X              = C.SDL_SCANCODE_X
	Key_Y              = C.SDL_SCANCODE_Y
	Key_Z              = C.SDL_SCANCODE_Z
	Key_1              = C.SDL_SCANCODE_1
	Key_2              = C.SDL_SCANCODE_2
	Key_3              = C.SDL_SCANCODE_3
	Key_4              = C.SDL_SCANCODE_4
	Key_5              = C.SDL_SCANCODE_5
	Key_6              = C.SDL_SCANCODE_6
	Key_7              = C.SDL_SCANCODE_7
	Key_8              = C.SDL_SCANCODE_8
	Key_9              = C.SDL_SCANCODE_9
	Key_0              = C.SDL_SCANCODE_0
	Key_RETURN         = C.SDL_SCANCODE_RETURN
	Key_ESCAPE         = C.SDL_SCANCODE_ESCAPE
	Key_BACKSPACE      = C.SDL_SCANCODE_BACKSPACE
	Key_TAB            = C.SDL_SCANCODE_TAB
	Key_SPACE          = C.SDL_SCANCODE_SPACE
	Key_MINUS          = C.SDL_SCANCODE_MINUS
	Key_EQUALS         = C.SDL_SCANCODE_EQUALS
	Key_LEFTBRACKET    = C.SDL_SCANCODE_LEFTBRACKET
	Key_RIGHTBRACKET   = C.SDL_SCANCODE_RIGHTBRACKET
	Key_BACKSLASH      = C.SDL_SCANCODE_BACKSLASH
	Key_NONUSHASH      = C.SDL_SCANCODE_NONUSHASH
	Key_SEMICOLON      = C.SDL_SCANCODE_SEMICOLON
	Key_APOSTROPHE     = C.SDL_SCANCODE_APOSTROPHE
	Key_GRAVE          = C.SDL_SCANCODE_GRAVE
	Key_COMMA          = C.SDL_SCANCODE_COMMA
	Key_PERIOD         = C.SDL_SCANCODE_PERIOD
	Key_SLASH          = C.SDL_SCANCODE_SLASH
	Key_CAPSLOCK       = C.SDL_SCANCODE_CAPSLOCK
	Key_F1             = C.SDL_SCANCODE_F1
	Key_F2             = C.SDL_SCANCODE_F2
	Key_F3             = C.SDL_SCANCODE_F3
	Key_F4             = C.SDL_SCANCODE_F4
	Key_F5             = C.SDL_SCANCODE_F5
	Key_F6             = C.SDL_SCANCODE_F6
	Key_F7             = C.SDL_SCANCODE_F7
	Key_F8             = C.SDL_SCANCODE_F8
	Key_F9             = C.SDL_SCANCODE_F9
	Key_F10            = C.SDL_SCANCODE_F10
	Key_F11            = C.SDL_SCANCODE_F11
	Key_F12            = C.SDL_SCANCODE_F12
	Key_PRINTSCREEN    = C.SDL_SCANCODE_PRINTSCREEN
	Key_SCROLLLOCK     = C.SDL_SCANCODE_SCROLLLOCK
	Key_PAUSE          = C.SDL_SCANCODE_PAUSE
	Key_INSERT         = C.SDL_SCANCODE_INSERT
	Key_HOME           = C.SDL_SCANCODE_HOME
	Key_PAGEUP         = C.SDL_SCANCODE_PAGEUP
	Key_DELETE         = C.SDL_SCANCODE_DELETE
	Key_END            = C.SDL_SCANCODE_END
	Key_PAGEDOWN       = C.SDL_SCANCODE_PAGEDOWN
	Key_RIGHT          = C.SDL_SCANCODE_RIGHT
	Key_LEFT           = C.SDL_SCANCODE_LEFT
	Key_DOWN           = C.SDL_SCANCODE_DOWN
	Key_UP             = C.SDL_SCANCODE_UP
	Key_NUMLOCKCLEAR   = C.SDL_SCANCODE_NUMLOCKCLEAR
	Key_KP_DIVIDE      = C.SDL_SCANCODE_KP_DIVIDE
	Key_KP_MULTIPLY    = C.SDL_SCANCODE_KP_MULTIPLY
	Key_KP_MINUS       = C.SDL_SCANCODE_KP_MINUS
	Key_KP_PLUS        = C.SDL_SCANCODE_KP_PLUS
	Key_KP_ENTER       = C.SDL_SCANCODE_KP_ENTER
	Key_KP_1           = C.SDL_SCANCODE_KP_1
	Key_KP_2           = C.SDL_SCANCODE_KP_2
	Key_KP_3           = C.SDL_SCANCODE_KP_3
	Key_KP_4           = C.SDL_SCANCODE_KP_4
	Key_KP_5           = C.SDL_SCANCODE_KP_5
	Key_KP_6           = C.SDL_SCANCODE_KP_6
	Key_KP_7           = C.SDL_SCANCODE_KP_7
	Key_KP_8           = C.SDL_SCANCODE_KP_8
	Key_KP_9           = C.SDL_SCANCODE_KP_9
	Key_KP_0           = C.SDL_SCANCODE_KP_0
	Key_KP_PERIOD      = C.SDL_SCANCODE_KP_PERIOD
	Key_NONUSBACKSLASH = C.SDL_SCANCODE_NONUSBACKSLASH
	Key_APPLICATION    = C.SDL_SCANCODE_APPLICATION
	Key_POWER          = C.SDL_SCANCODE_POWER
	Key_KP_EQUALS      = C.SDL_SCANCODE_KP_EQUALS
	Key_F13            = C.SDL_SCANCODE_F13
	Key_F14            = C.SDL_SCANCODE_F14
	Key_F15            = C.SDL_SCANCODE_F15
	Key_F16            = C.SDL_SCANCODE_F16
	Key_F17            = C.SDL_SCANCODE_F17
	Key_F18            = C.SDL_SCANCODE_F18
	Key_F19            = C.SDL_SCANCODE_F19
	Key_F20            = C.SDL_SCANCODE_F20
	Key_F21            = C.SDL_SCANCODE_F21
	Key_F22            = C.SDL_SCANCODE_F22
	Key_F23            = C.SDL_SCANCODE_F23
	Key_F24            = C.SDL_SCANCODE_F24
	Key_EXECUTE        = C.SDL_SCANCODE_EXECUTE
	Key_HELP           = C.SDL_SCANCODE_HELP
	Key_MENU           = C.SDL_SCANCODE_MENU
	Key_SELECT         = C.SDL_SCANCODE_SELECT
	Key_STOP           = C.SDL_SCANCODE_STOP
	Key_AGAIN          = C.SDL_SCANCODE_AGAIN
	Key_UNDO           = C.SDL_SCANCODE_UNDO
	Key_CUT            = C.SDL_SCANCODE_CUT
	Key_COPY           = C.SDL_SCANCODE_COPY
	Key_PASTE          = C.SDL_SCANCODE_PASTE
	Key_FIND           = C.SDL_SCANCODE_FIND
	Key_MUTE           = C.SDL_SCANCODE_MUTE
	Key_VOLUMEUP       = C.SDL_SCANCODE_VOLUMEUP
	Key_VOLUMEDOWN     = C.SDL_SCANCODE_VOLUMEDOWN
	//Key_LOCKINGCAPSLOCK    = C.SDL_SCANCODE_LOCKINGCAPSLOCK
	//Key_LOCKINGNUMLOCK     = C.SDL_SCANCODE_LOCKINGNUMLOCK
	//Key_LOCKINGSCROLLLOCK  = C.SDL_SCANCODE_LOCKINGSCROLLLOCK
	Key_KP_COMMA           = C.SDL_SCANCODE_KP_COMMA
	Key_KP_EQUALSAS400     = C.SDL_SCANCODE_KP_EQUALSAS400
	Key_INTERNATIONAL1     = C.SDL_SCANCODE_INTERNATIONAL1
	Key_INTERNATIONAL2     = C.SDL_SCANCODE_INTERNATIONAL2
	Key_INTERNATIONAL3     = C.SDL_SCANCODE_INTERNATIONAL3
	Key_INTERNATIONAL4     = C.SDL_SCANCODE_INTERNATIONAL4
	Key_INTERNATIONAL5     = C.SDL_SCANCODE_INTERNATIONAL5
	Key_INTERNATIONAL6     = C.SDL_SCANCODE_INTERNATIONAL6
	Key_INTERNATIONAL7     = C.SDL_SCANCODE_INTERNATIONAL7
	Key_INTERNATIONAL8     = C.SDL_SCANCODE_INTERNATIONAL8
	Key_INTERNATIONAL9     = C.SDL_SCANCODE_INTERNATIONAL9
	Key_LANG1              = C.SDL_SCANCODE_LANG1
	Key_LANG2              = C.SDL_SCANCODE_LANG2
	Key_LANG3              = C.SDL_SCANCODE_LANG3
	Key_LANG4              = C.SDL_SCANCODE_LANG4
	Key_LANG5              = C.SDL_SCANCODE_LANG5
	Key_LANG6              = C.SDL_SCANCODE_LANG6
	Key_LANG7              = C.SDL_SCANCODE_LANG7
	Key_LANG8              = C.SDL_SCANCODE_LANG8
	Key_LANG9              = C.SDL_SCANCODE_LANG9
	Key_ALTERASE           = C.SDL_SCANCODE_ALTERASE
	Key_SYSREQ             = C.SDL_SCANCODE_SYSREQ
	Key_CANCEL             = C.SDL_SCANCODE_CANCEL
	Key_CLEAR              = C.SDL_SCANCODE_CLEAR
	Key_PRIOR              = C.SDL_SCANCODE_PRIOR
	Key_RETURN2            = C.SDL_SCANCODE_RETURN2
	Key_SEPARATOR          = C.SDL_SCANCODE_SEPARATOR
	Key_OUT                = C.SDL_SCANCODE_OUT
	Key_OPER               = C.SDL_SCANCODE_OPER
	Key_CLEARAGAIN         = C.SDL_SCANCODE_CLEARAGAIN
	Key_CRSEL              = C.SDL_SCANCODE_CRSEL
	Key_EXSEL              = C.SDL_SCANCODE_EXSEL
	Key_KP_00              = C.SDL_SCANCODE_KP_00
	Key_KP_000             = C.SDL_SCANCODE_KP_000
	Key_THOUSANDSSEPARATOR = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	Key_DECIMALSEPARATOR   = C.SDL_SCANCODE_DECIMALSEPARATOR
	Key_CURRENCYUNIT       = C.SDL_SCANCODE_CURRENCYUNIT
	Key_CURRENCYSUBUNIT    = C.SDL_SCANCODE_CURRENCYSUBUNIT
	Key_KP_LEFTPAREN       = C.SDL_SCANCODE_KP_LEFTPAREN
	Key_KP_RIGHTPAREN      = C.SDL_SCANCODE_KP_RIGHTPAREN
	Key_KP_LEFTBRACE       = C.SDL_SCANCODE_KP_LEFTBRACE
	Key_KP_RIGHTBRACE      = C.SDL_SCANCODE_KP_RIGHTBRACE
	Key_KP_TAB             = C.SDL_SCANCODE_KP_TAB
	Key_KP_BACKSPACE       = C.SDL_SCANCODE_KP_BACKSPACE
	Key_KP_A               = C.SDL_SCANCODE_KP_A
	Key_KP_B               = C.SDL_SCANCODE_KP_B
	Key_KP_C               = C.SDL_SCANCODE_KP_C
	Key_KP_D               = C.SDL_SCANCODE_KP_D
	Key_KP_E               = C.SDL_SCANCODE_KP_E
	Key_KP_F               = C.SDL_SCANCODE_KP_F
	Key_KP_XOR             = C.SDL_SCANCODE_KP_XOR
	Key_KP_POWER           = C.SDL_SCANCODE_KP_POWER
	Key_KP_PERCENT         = C.SDL_SCANCODE_KP_PERCENT
	Key_KP_LESS            = C.SDL_SCANCODE_KP_LESS
	Key_KP_GREATER         = C.SDL_SCANCODE_KP_GREATER
	Key_KP_AMPERSAND       = C.SDL_SCANCODE_KP_AMPERSAND
	Key_KP_DBLAMPERSAND    = C.SDL_SCANCODE_KP_DBLAMPERSAND
	Key_KP_VERTICALBAR     = C.SDL_SCANCODE_KP_VERTICALBAR
	Key_KP_DBLVERTICALBAR  = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	Key_KP_COLON           = C.SDL_SCANCODE_KP_COLON
	Key_KP_HASH            = C.SDL_SCANCODE_KP_HASH
	Key_KP_SPACE           = C.SDL_SCANCODE_KP_SPACE
	Key_KP_AT              = C.SDL_SCANCODE_KP_AT
	Key_KP_EXCLAM          = C.SDL_SCANCODE_KP_EXCLAM
	Key_KP_MEMSTORE        = C.SDL_SCANCODE_KP_MEMSTORE
	Key_KP_MEMRECALL       = C.SDL_SCANCODE_KP_MEMRECALL
	Key_KP_MEMCLEAR        = C.SDL_SCANCODE_KP_MEMCLEAR
	Key_KP_MEMADD          = C.SDL_SCANCODE_KP_MEMADD
	Key_KP_MEMSUBTRACT     = C.SDL_SCANCODE_KP_MEMSUBTRACT
	Key_KP_MEMMULTIPLY     = C.SDL_SCANCODE_KP_MEMMULTIPLY
	Key_KP_MEMDIVIDE       = C.SDL_SCANCODE_KP_MEMDIVIDE
	Key_KP_PLUSMINUS       = C.SDL_SCANCODE_KP_PLUSMINUS
	Key_KP_CLEAR           = C.SDL_SCANCODE_KP_CLEAR
	Key_KP_CLEARENTRY      = C.SDL_SCANCODE_KP_CLEARENTRY
	Key_KP_BINARY          = C.SDL_SCANCODE_KP_BINARY
	Key_KP_OCTAL           = C.SDL_SCANCODE_KP_OCTAL
	Key_KP_DECIMAL         = C.SDL_SCANCODE_KP_DECIMAL
	Key_KP_HEXADECIMAL     = C.SDL_SCANCODE_KP_HEXADECIMAL
	Key_LCTRL              = C.SDL_SCANCODE_LCTRL
	Key_LSHIFT             = C.SDL_SCANCODE_LSHIFT
	Key_LALT               = C.SDL_SCANCODE_LALT
	Key_LGUI               = C.SDL_SCANCODE_LGUI
	Key_RCTRL              = C.SDL_SCANCODE_RCTRL
	Key_RSHIFT             = C.SDL_SCANCODE_RSHIFT
	Key_RALT               = C.SDL_SCANCODE_RALT
	Key_RGUI               = C.SDL_SCANCODE_RGUI
	Key_MODE               = C.SDL_SCANCODE_MODE
	Key_AUDIONEXT          = C.SDL_SCANCODE_AUDIONEXT
	Key_AUDIOPREV          = C.SDL_SCANCODE_AUDIOPREV
	Key_AUDIOSTOP          = C.SDL_SCANCODE_AUDIOSTOP
	Key_AUDIOPLAY          = C.SDL_SCANCODE_AUDIOPLAY
	Key_AUDIOMUTE          = C.SDL_SCANCODE_AUDIOMUTE
	Key_MEDIASELECT        = C.SDL_SCANCODE_MEDIASELECT
	Key_WWW                = C.SDL_SCANCODE_WWW
	Key_MAIL               = C.SDL_SCANCODE_MAIL
	Key_CALCULATOR         = C.SDL_SCANCODE_CALCULATOR
	Key_COMPUTER           = C.SDL_SCANCODE_COMPUTER
	Key_AC_SEARCH          = C.SDL_SCANCODE_AC_SEARCH
	Key_AC_HOME            = C.SDL_SCANCODE_AC_HOME
	Key_AC_BACK            = C.SDL_SCANCODE_AC_BACK
	Key_AC_FORWARD         = C.SDL_SCANCODE_AC_FORWARD
	Key_AC_STOP            = C.SDL_SCANCODE_AC_STOP
	Key_AC_REFRESH         = C.SDL_SCANCODE_AC_REFRESH
	Key_AC_BOOKMARKS       = C.SDL_SCANCODE_AC_BOOKMARKS
	Key_BRIGHTNESSDOWN     = C.SDL_SCANCODE_BRIGHTNESSDOWN
	Key_BRIGHTNESSUP       = C.SDL_SCANCODE_BRIGHTNESSUP
	Key_DISPLAYSWITCH      = C.SDL_SCANCODE_DISPLAYSWITCH
	Key_KBDILLUMTOGGLE     = C.SDL_SCANCODE_KBDILLUMTOGGLE
	Key_KBDILLUMDOWN       = C.SDL_SCANCODE_KBDILLUMDOWN
	Key_KBDILLUMUP         = C.SDL_SCANCODE_KBDILLUMUP
	Key_EJECT              = C.SDL_SCANCODE_EJECT
	Key_SLEEP              = C.SDL_SCANCODE_SLEEP
	Key_APP1               = C.SDL_SCANCODE_APP1
	Key_APP2               = C.SDL_SCANCODE_APP2
)
