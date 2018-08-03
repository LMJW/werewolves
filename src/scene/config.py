# Game settings

# Text color of menu
_DEFAULT_TEXT_COLOR = '#678900'
_DEFAULT_SELECTED_TEXT_COLOR = '#50d661'

_MAX_NUMBER_CHARACTERS = 24

_multiplayer = [i for i in range(_MAX_NUMBER_CHARACTERS)]
_singleplayer = [0, 1]

_GAME_CHARACTERS = {
    "Villager": _multiplayer,
    "Werewolf": _multiplayer,
    "Minion": _singleplayer,
    "Mason": _multiplayer,
    "Seer": _singleplayer,
    "Robber": _singleplayer,
    "Trouble maker": _singleplayer,
    "Drunk": _singleplayer,
    "Insomniac": _singleplayer,
    "Hunter": _singleplayer,
    "Tanner": _multiplayer
}
