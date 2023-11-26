package main

type KeysLeftRight int
type KeysUpDown int
type Block int
type Listener int

const (
	// object types
	OBJECT_TYPE_SPIKE     Block = 0
	OBJECT_TYPE_BLOCK     Block = 1
	OBJECT_TYPE_CHERRY    Block = 2
	OBJECT_TYPE_SAVEPOINT Block = 3

	OBJECT_TYPE_STARTFLAG  Block = 5
	OBJECT_TYPE_FINISHWARP Block = 6
	OBJECT_TYPE_MINISPIKE  Block = 7
	OBJECT_TYPE_TRIGGER    Block = 8

	OBJECT_TYPE_SPINNER Block = 18

	OBJECT_TYPE_WATER1         Block = 20
	OBJECT_TYPE_WATER2         Block = 21
	OBJECT_TYPE_PLATFORM       Block = 22
	OBJECT_TYPE_JUMP_REFRESHER Block = 23
	OBJECT_TYPE_MINIBLOCK      Block = 24
	OBJECT_TYPE_CONVEYOR       Block = 25
	OBJECT_TYPE_GRAVITY_UP     Block = 26
	OBJECT_TYPE_GRAVITY_DOWN   Block = 27
	OBJECT_TYPE_CANNON         Block = 28

	OBJECT_TYPE_TELEKID_ON         Block = 35
	OBJECT_TYPE_TELEKID_OFF        Block = 36
	OBJECT_TYPE_3JUMPSTAR          Block = 37
	OBJECT_TYPE_2JUMPSTAR          Block = 38
	OBJECT_TYPE_1JUMPSTAR          Block = 39
	OBJECT_TYPE_LOWSPEED_FIELD     Block = 40
	OBJECT_TYPE_HIGHSPEED_FIELD    Block = 41
	OBJECT_TYPE_LOWGRAVITY_FIELD   Block = 42
	OBJECT_TYPE_HIGHGRAVITY_FIELD  Block = 43
	OBJECT_TYPE_ONEWAY_WALL        Block = 44
	OBJECT_TYPE_TELEPORT_ON        Block = 45
	OBJECT_TYPE_TELEPORT_OFF       Block = 46
	OBJECT_TYPE_TELEPORT_REFRESHER Block = 47
	OBJECT_TYPE_SPRING             Block = 48
	OBJECT_TYPE_ARROW              Block = 49
	OBJECT_TYPE_WEAPON_ARROW       Block = 50
	OBJECT_TYPE_COIN               Block = 51
	OBJECT_TYPE_COIN_BLOCK         Block = 52

	OBJECT_TYPE_SHURIKEN   Block = 71
	OBJECT_TYPE_BUTTON     Block = 72
	OBJECT_TYPE_CRUSHER    Block = 73
	OBJECT_TYPE_BACKGROUND Block = 74

	OBJECT_TYPE_CANNON_BALL Block = 78
	OBJECT_TYPE_REFLECTOR   Block = 79

	OBJECT_TYPE_LASER              Block = 83
	OBJECT_TYPE_ACCELERATION_FIELD Block = 84
	OBJECT_TYPE_TREE               Block = 85
	OBJECT_TYPE_CAUTION_SIGN       Block = 86
	OBJECT_TYPE_TEXT_SIGN          Block = 87
	OBJECT_TYPE_EFFECT_TRIGGER     Block = 88
	OBJECT_TYPE_EFFECT_REMOVER     Block = 89
	OBJECT_TYPE_OBJECT_ACTIVATOR   Block = 90
	OBJECT_TYPE_BALL               Block = 91
	OBJECT_TYPE_SWORD_ACTIVATE     Block = 92
	OBJECT_TYPE_SWORD_DEACTIVATE   Block = 93
	OBJECT_TYPE_TELEPORTER_IN      Block = 94
	OBJECT_TYPE_TELEPORTER_OUT     Block = 95
	OBJECT_TYPE_KILLER_BLOCK       Block = 96

	OBJECT_TYPE_SABER       Block = 98
	OBJECT_TYPE_PARACHUTE   Block = 99
	OBJECT_TYPE_SHIELD      Block = 100
	OBJECT_TYPE_PUSH_BLOCK  Block = 101
	OBJECT_TYPE_WALKER_BIRD Block = 102

	OBJECT_TYPE_PARACHUTE_REMOVER Block = 104
	OBJECT_TYPE_BIG_KID           Block = 105
	OBJECT_TYPE_NORMAL_KID        Block = 106
	OBJECT_TYPE_SMALL_KID         Block = 107
	OBJECT_TYPE_BLUE_COIN         Block = 108
	OBJECT_TYPE_BLUE_COIN_BLOCK   Block = 109
	OBJECT_TYPE_LIGHT_POST        Block = 110
	OBJECT_TYPE_CRUSHER_STATUE    Block = 111
	OBJECT_TYPE_MUSHROOM          Block = 112
	OBJECT_TYPE_HEART             Block = 113
	OBJECT_TYPE_FLOWER            Block = 114
	OBJECT_TYPE_SPARKLE           Block = 115
	OBJECT_TYPE_TORCH             Block = 116
	OBJECT_TYPE_CAMPFIRE          Block = 117
	OBJECT_TYPE_COUCH             Block = 119
	OBJECT_TYPE_BULLET            Block = 120
	OBJECT_TYPE_BUBBLE_ON         Block = 121
	OBJECT_TYPE_BUBBLE_OFF        Block = 122
	OBJECT_TYPE_LIGHT             Block = 123
	OBJECT_TYPE_TARGET            Block = 124
	OBJECT_TYPE_GHOST             Block = 125
	OBJECT_TYPE_BURST_BLOCK       Block = 126
	OBJECT_TYPE_SIGN              Block = 131
	OBJECT_TYPE_ANCHOR_FIELD      Block = 132

	// on-event types
	EVENT_ON_CREATE              Listener = 1
	EVENT_ON_METRONOME_TICK      Listener = 17
	EVENT_ON_TIMER               Listener = 2
	EVENT_WHEN_SHOT              Listener = 15
	EVENT_WHEN_HIT_BY_SWORD      Listener = 22
	EVENT_ON_PLAYER_JUMP         Listener = 13
	EVENT_ON_PLAYER_SHOOT        Listener = 14
	EVENT_ON_LEFT_RIGHT_PRESSED  Listener = 24
	EVENT_ON_UP_DOWN_PRESSED     Listener = 28
	EVENT_ON_VINE_JUMP           Listener = 20
	EVENT_ON_CANNON_BALL_BOUNCE  Listener = 23
	EVENT_ON_SPRUNG              Listener = 21
	EVENT_ON_COINS_COLLECTED     Listener = 18
	EVENT_ON_A_COIN_COLLECTED    Listener = 26
	EVENT_ON_BLUE_COIN_COLLECTED Listener = 27
	EVENT_ON_SHIELD_LOSS         Listener = 29
	EVENT_ON_TRIGGER             Listener = 16
	EVENT_ON_TOUCH_ACTIVATOR     Listener = 19
	EVENT_ON_OBJECT_COLLISION    Listener = 25
	EVENT_ON_TOUCHED_EXPLOSION   Listener = 30

	// trigger-event types
	EVENT_DESTROY            = 103
	EVENT_TOGGLE_COLLISIONS  = 107
	EVENT_START_MOVING       = 101
	EVENT_MOVE_SET_DISTANCE  = 105
	EVENT_FOLLOW_PATH        = 116
	EVENT_MOVE_TO_POSITION   = 122
	EVENT_MOVE_TO_PLAYER     = 123
	EVENT_FOLLOW_PLAYER_AXIS = 120
	EVENT_STOP_MOVING        = 106
	EVENT_SET_SPEED          = 124
	EVENT_SET_ANGLE          = 125
	EVENT_SET_ACCELERATION   = 113
	EVENT_PLAY_SOUND         = 104
	EVENT_PLAY_MUSIC         = 119
	EVENT_SET_MUSIC_PITCH    = 118
	EVENT_CAMERA_FOLLOW      = 117
	EVENT_ACTIVATE_EFFECT    = 130
	EVENT_REMOVE_EFFECTS     = 131
	EVENT_SET_TIMER          = 102
	EVENT_STOP_TIMER         = 114
	EVENT_ACTIVATE_TRIGGER   = 110
	EVENT_SET_RANDOM_SPEED   = 134
	EVENT_SET_RANDOM_ANGLE   = 133

	EVENT_ROTATE_90_DEGREES = 115

	EVENT_SET_CHERRY_COLOR         = 126
	EVENT_SET_BARRAGE_BULLET_COLOR = 127
	EVENT_TOGGLE_BOUNCE            = 137

	EVENT_SET_ROTATION        = 132
	EVENT_SET_ROTATE_SPEED    = 111
	EVENT_SET_RANDOM_ROTATION = 136

	EVENT_SET_WALKER_WALK_SPEED = 128

	EVENT_RADIUS_SHIFT_SPEED = 121

	EVENT_FIRE_CANNON            = 109
	EVENT_TOGGLE_CANNON_TRACKING = 129
	EVENT_FIRE_RANDOM_SLOT       = 135

	EVENT_SET_CONVEYOR_SPEED = 112

	EVENT_ACTIVATE_BURST_BLOCK = 138
	EVENT_CHANGE_SIGN_VISIBLE  = 139
	CHERRY_COLOR_RED           = 0
	CHERRY_COLOR_ORANGE        = 1
	CHERRY_COLOR_YELLOW        = 2
	CHERRY_COLOR_GREEN         = 3
	CHERRY_COLOR_LIME          = 4
	CHERRY_COLOR_BLUE          = 5
	CHERRY_COLOR_PURPLE        = 6
	CHERRY_COLOR_SILVER        = 7
	CHERRY_COLOR_GRAY          = 8
	CHERRY_COLOR_PURE_BLACK    = 9
	CHERRY_COLOR_PURE_WHITE    = 10
	CHERRY_COLOR_RANDOM        = 11

	GRAV_TYPE_UP   = 0
	GRAV_TYPE_DOWN = 1

	// sounds
	DOWNED_G                       = "0.1873"
	DOWNED_G_SHARP                 = "0.1984"
	DOWNED_A                       = "0.2102"
	DOWNED_A_SHARP                 = "0.2227"
	DOWNED_B                       = "0.236"
	DOWNED_C                       = "0.25"
	DOWNED_C_SHARP                 = "0.2649"
	DOWNED_D                       = "0.2806"
	DOWNED_D_SHARP                 = "0.2973"
	DOWNED_E                       = "0.315"
	DOWNED_F                       = "0.3337"
	DOWNED_F_SHARP                 = "0.3536"
	BASIC_G                        = "0.3746" // ド
	BASIC_G_SHARP                  = "0.3968" // ド#
	BASIC_A                        = "0.4204" // レ
	BASIC_A_SHARP                  = "0.4454" // レ#
	BASIC_B                        = "0.4719" // ミ
	BASIC_C                        = "0.5"    // ファ
	BASIC_C_SHARP                  = "0.5297" // ファ#
	BASIC_D                        = "0.5612" // ソ
	BASIC_D_SHARP                  = "0.5946" // ソ#
	BASIC_E                        = "0.63"   // ラ
	BASIC_F                        = "0.6674" // シ
	BASIC_F_SHARP                  = "0.3536" // シ#
	ONE_UPED_G                     = "0.7491" // ド
	ONE_UPED_G_SHARP               = "0.7937"
	ONE_UPED_A                     = "0.8409"
	ONE_UPED_A_SHARP               = "0.8909"
	ONE_UPED_B                     = "0.9439"
	ONE_UPED_C                     = "1"
	ONE_UPED_C_SHARP               = "1.0595"
	ONE_UPED_D                     = "1.1225"
	ONE_UPED_D_SHARP               = "1.1892"
	ONE_UPED_E                     = "1.2599"
	ONE_UPED_F                     = "1.3348"
	ONE_UPED_F_SHARP               = "1.4142"
	LEFT             KeysLeftRight = 0
	RIGHT            KeysLeftRight = 1
	UP               KeysUpDown    = 0
	DOWN             KeysUpDown    = 1
)
