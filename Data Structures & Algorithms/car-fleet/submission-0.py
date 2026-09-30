class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        fleets = []

        pos_sp = sorted(zip(position, speed), key = lambda duo: duo[0])

        for pos, sp in pos_sp:
            while len(fleets) > 0:
                pos_prev = fleets[-1][0]
                sp_prev = fleets[-1][1]

                time_left = (target - pos) / sp
                end_prev = pos_prev + sp_prev * time_left

                if end_prev >= target:
                    fleets.pop()
                else:
                    break
                
            fleets.append((pos, sp))

        return len(fleets)
        