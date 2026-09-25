class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        right = []
        left = []

        for a in asteroids:
            if a > 0:
                right.append(a)
            else:
                broken = False

                while len(right) > 0:
                    if right[-1] < -a:
                        right.pop()
                    elif right[-1] == -a:
                        right.pop()
                        broken = True
                        break
                    else:
                        broken = True
                        break

                if not broken:
                    left.append(a)

        return left + right
        