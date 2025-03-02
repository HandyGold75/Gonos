package Gonos

import "strconv"

// Short for `zp.AudioIn.GetLineInLevel`.
func (zp *ZonePlayer) GetLineInLevel() (int, int, error) {
	res, err := zp.AudioIn.GetLineInLevel()
	if err != nil {
		return 0, 0, err
	}
	return res.CurrentLeftLineInLevel, res.CurrentRightLineInLevel, nil
}

// Short for `zp.AudioIn.GetLineInLevel` (Left).
func (zp *ZonePlayer) GetLineInLevelLeft() (int, error) {
	res, err := zp.AudioIn.GetLineInLevel()
	if err != nil {
		return 0, err
	}
	return res.CurrentLeftLineInLevel, nil
}

// Short for `zp.AudioIn.GetLineInLevel` (Right).
func (zp *ZonePlayer) GetLineInLevelRight() (int, error) {
	res, err := zp.AudioIn.GetLineInLevel()
	if err != nil {
		return 0, err
	}
	return res.CurrentRightLineInLevel, nil
}

// Short for `zp.AudioIn.SetLineInLevel`.
func (zp *ZonePlayer) SetLineInLevel(volume int) error {
	return zp.AudioIn.SetLineInLevel(volume, volume)
}

// Short for `zp.AudioIn.SetLineInLevel` (Left).
func (zp *ZonePlayer) SetLineInLevelLeft(volume int) error {
	_, err := zp.AudioIn.Send("SetLineInLevel", "<DesiredLeftLineInLevel>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredLeftLineInLevel>", "")
	return err
}

// Short for `zp.AudioIn.SetLineInLevel` (Right).
func (zp *ZonePlayer) SetLineInLevelRight(volume int) error {
	_, err := zp.AudioIn.Send("SetLineInLevel", "<DesiredRightLineInLevel>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredRightLineInLevel>", "")
	return err
}
