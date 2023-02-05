import re
from pathlib import Path

import numpy as np


MOVE_PAT = re.compile(r'\d+|[A-Z]')


def parse_board(lines):
    rows = []
    lines = lines.split('\n')

    max_length = max(len(line) for line in lines)

    for line in lines:
        row = [' .#'.index(c) for c in line]
        row.extend([0 for _ in range(max_length - len(row))])
        rows.append(row)

    return np.array(rows)


def parse_route(line):
    line = line.strip()

    route = []

    for step in MOVE_PAT.findall(line):
        try:
            step = int(step)
            route.append(('move', int(step)))
        except ValueError:
            route.append(('turn', step))

    return route


def part2(s: str):
    board_lines, route_line = s.split('\n\n')

    full_board = parse_board(board_lines)
    route = parse_route(route_line)

    ln = 50
    origins = {
        'top': (0 * ln, 1 * ln),
        'right': (0 * ln, 2 * ln),

        'front': (1 * ln, 1 * ln),

        'left': (2 * ln, 0 * ln),
        'bot': (2 * ln, 1 * ln),

        'back': (3 * ln, 0 * ln),
    }

    subboards = {key: full_board[ro:ro + ln, co:co + ln] for key, (ro, co) in origins.items()}

    right = (0, 1)
    down = (1, 0)
    left = (0, -1)
    up = (-1, 0)

    facings = (right, down, left, up)

    def rotate_right(facing):
        i = facings.index(facing)
        return facings[(i + 1) % 4]

    def rotate_left(facing):
        i = facings.index(facing)
        return facings[(i - 1) % 4]

    def teleport(pos, facing, side):
        r, c = pos

        if (side, facing) == ('top', up):
            new_side, new_facing = 'back', right
            r = c
            c = 0

        elif (side, facing) == ('top', down):
            new_side, new_facing = 'front', down
            r = 0

        elif (side, facing) == ('top', left):
            new_side, new_facing = 'left', right
            r = ln - 1 - r

        elif (side, facing) == ('top', right):
            new_side, new_facing = 'right', right
            c = 0

        elif (side, facing) == ('bot', up):
            new_side, new_facing = 'front', up
            r = ln - 1

        elif (side, facing) == ('bot', down):
            new_side, new_facing = 'back', left
            r = c
            c = ln - 1

        elif (side, facing) == ('bot', left):
            new_side, new_facing = 'left', left
            c = ln - 1

        elif (side, facing) == ('bot', right):
            new_side, new_facing = 'right', left
            r = ln - 1 - r

        elif (side, facing) == ('left', up):
            new_side, new_facing = 'front', right
            r = c
            c = 0

        elif (side, facing) == ('left', down):
            new_side, new_facing = 'back', down
            r = 0

        elif (side, facing) == ('left', left):
            new_side, new_facing = 'top', right
            r = ln - 1 - r

        elif (side, facing) == ('left', right):
            new_side, new_facing = 'bot', right
            c = 0

        elif (side, facing) == ('right', up):
            new_side, new_facing = 'back', up
            r = ln - 1

        elif (side, facing) == ('right', down):
            new_side, new_facing = 'front', left
            r = c
            c = ln - 1

        elif (side, facing) == ('right', left):
            new_side, new_facing = 'top', left
            c = ln - 1

        elif (side, facing) == ('right', right):
            new_side, new_facing = 'bot', left
            r = ln - 1 - r

        elif (side, facing) == ('front', up):
            new_side, new_facing = 'top', up
            r = ln - 1

        elif (side, facing) == ('front', down):
            new_side, new_facing = 'bot', down
            r = 0

        elif (side, facing) == ('front', left):
            new_side, new_facing = 'left', down
            c = r
            r = 0

        elif (side, facing) == ('front', right):
            new_side, new_facing = 'right', up
            c = r
            r = ln - 1

        elif (side, facing) == ('back', up):
            new_side, new_facing = 'left', up
            r = ln - 1

        elif (side, facing) == ('back', down):
            new_side, new_facing = 'right', down
            r = 0

        elif (side, facing) == ('back', left):
            new_side, new_facing = 'top', down
            c = r
            r = 0

        elif (side, facing) == ('back', right):
            new_side, new_facing = 'bot', up
            c = r
            r = ln - 1

        new_pos = (r, c)

        return new_pos, new_facing, new_side

    def find_new_pos(pos, facing, side, *, max_n):
        accepted = (pos, facing, side)
        board = subboards[side]

        for n in range(max_n):
            dr, dc = facing
            r, c = pos
            new_side = side

            tr = r + dr
            tc = c + dc

            if tr in (-1, ln) or tc in (-1, ln):
                (tr, tc), (dr, dc), new_side = teleport((r, c), (dr, dc), side)
                board = subboards[new_side]

            if board[tr, tc] == 1:  # open
                pos = (tr, tc)
                facing = (dr, dc)
                side = new_side

                accepted = (pos, facing, side)

            elif board[tr, tc] == 2:  # rock
                break

            elif board[tr, tc] == 0:
                raise ValueError('This should no longer happen')

        return accepted

    side = 'top'
    pos = 0, 0
    facing = facings[0]

    for i, instruction in enumerate(route):
        match instruction:
            case 'move', n:
                (pos, facing, side) = find_new_pos(pos, facing, side, max_n=n)
            case 'turn', 'L':
                facing = rotate_left(facing)
            case 'turn', 'R':
                facing = rotate_right(facing)

    ro, rc = origins[side]
    r, c = pos

    return 1000 * (ro + r + 1) + 4 * (rc + c + 1) + facings.index(facing)


if __name__ == '__main__':
    DATA = Path(__file__).with_name('input.txt').read_text()

    print(part2(DATA))
