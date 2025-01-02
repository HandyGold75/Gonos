package Helper

import "strconv"

func (h *Helper) GetLineInLevel() (int, int, error) {
	res, err := h.audioIn.GetLineInLevel()
	if err != nil {
		return 0, 0, err
	}
	return res.CurrentLeftLineInLevel, res.CurrentRightLineInLevel, nil
}

func (h *Helper) GetLineInLevelLeft() (int, error) {
	res, err := h.audioIn.GetLineInLevel()
	if err != nil {
		return 0, err
	}
	return res.CurrentLeftLineInLevel, nil
}

func (h *Helper) GetLineInLevelRight() (int, error) {
	res, err := h.audioIn.GetLineInLevel()
	if err != nil {
		return 0, err
	}
	return res.CurrentRightLineInLevel, nil
}

func (h *Helper) SetLineInLevel(volume int) error {
	return h.audioIn.SetLineInLevel(volume, volume)
}

func (h *Helper) SetLineInLevelLeft(volume int) error {
	_, err := h.audioIn.Send("SetLineInLevel", "<DesiredLeftLineInLevel>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredLeftLineInLevel>", "")
	return err
}

func (h *Helper) SetLineInLevelRight(volume int) error {
	_, err := h.audioIn.Send("SetLineInLevel", "<DesiredRightLineInLevel>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredRightLineInLevel>", "")
	return err
}
