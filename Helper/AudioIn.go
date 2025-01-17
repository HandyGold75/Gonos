package Helper

import "strconv"

// Short for `zp.AudioIn.GetLineInLevel`.
func (h *Helper) GetLineInLevel() (int, int, error) {
	res, err := h.audioIn.GetLineInLevel()
	if err != nil {
		return 0, 0, err
	}
	return res.CurrentLeftLineInLevel, res.CurrentRightLineInLevel, nil
}

// Short for `zp.AudioIn.GetLineInLevel` (Left).
func (h *Helper) GetLineInLevelLeft() (int, error) {
	res, err := h.audioIn.GetLineInLevel()
	if err != nil {
		return 0, err
	}
	return res.CurrentLeftLineInLevel, nil
}

// Short for `zp.AudioIn.GetLineInLevel` (Right).
func (h *Helper) GetLineInLevelRight() (int, error) {
	res, err := h.audioIn.GetLineInLevel()
	if err != nil {
		return 0, err
	}
	return res.CurrentRightLineInLevel, nil
}

// Short for `zp.AudioIn.SetLineInLevel`.
func (h *Helper) SetLineInLevel(volume int) error {
	return h.audioIn.SetLineInLevel(volume, volume)
}

// Short for `zp.AudioIn.SetLineInLevel` (Left).
func (h *Helper) SetLineInLevelLeft(volume int) error {
	_, err := h.audioIn.Send("SetLineInLevel", "<DesiredLeftLineInLevel>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredLeftLineInLevel>", "")
	return err
}

// Short for `zp.AudioIn.SetLineInLevel` (Right).
func (h *Helper) SetLineInLevelRight(volume int) error {
	_, err := h.audioIn.Send("SetLineInLevel", "<DesiredRightLineInLevel>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredRightLineInLevel>", "")
	return err
}
