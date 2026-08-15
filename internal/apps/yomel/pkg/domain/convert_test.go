package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

func TestConvert(t *testing.T) {
	ctrlModel := model.ControlModel{
		IsGen:                true,
		IsLog:                testutil.Ptr(true),
		LogFilter:            "para-log-filter",
		ErrLogFilter:         "para-err-filter",
		Title:                "Test Title",
		IsVersion:            false,
		IsHelp:               false,
		IsDirect:             false,
		IsLiveStdout:         true,
		IsLiveStderr:         true,
		ColorStr:             "red>blue",
		BgColorStr:           "green>black",
		CommentColorStr:      "yellow",
		TitleColorStr:        "blue",
		TitleBgColorStr:      "white",
		TitleCommentColorStr: "gray",
	}

	stModels := []model.StageModel{
		{
			No:   1,
			Desc: "Stage 1",
			Cmd:  "echo",
			CmdOps: []model.OptParam{
				{
					Index:  1,
					OptStr: "v",
					Param: model.ParamType{
						Str:       testutil.Ptr("val1"),
						QuoteType: argtables.DoubleQuote,
						Comment:   "OptComment",
					},
				},
				{
					Index:  2,
					OptStr: "n",
					Param: model.ParamType{
						Str:       nil,
						QuoteType: argtables.NoQuote,
						Comment:   "NoValueComment",
					},
				},
			},
			CmdLops: []model.OptParam{
				{
					Index:  3,
					OptStr: "long",
					Param: model.ParamType{
						Str:       testutil.Ptr("val2"),
						QuoteType: argtables.SingleQuote,
					},
				},
			},
			CmdArgs: []model.ArgParam{
				{
					Index: 4,
					Param: model.ParamType{
						Str:       testutil.Ptr("arg1"),
						QuoteType: argtables.NoQuote,
						Comment:   "ArgComment",
					},
				},
				{
					Index: 5,
					Param: model.ParamType{
						Str: nil,
					},
				},
			},
			Svc:             testutil.Ptr("service"),
			SvcOps:          []model.OptParam{},
			SvcLops:         []model.OptParam{},
			SvcArgs:         []model.ArgParam{},
			Act:             testutil.Ptr("action"),
			ActOps:          []model.OptParam{},
			ActLops:         []model.OptParam{},
			ActArgs:         []model.ArgParam{},
			IsLog:           testutil.Ptr(false),
			LogFilter:       "st-log-filter",
			ErrLogFilter:    "st-err-filter",
			ColorStr:        "red",
			BgColorStr:      "green",
			CommentColorStr: "blue",
		},
	}

	yomelToml := toml.LogConfig{
		Color: toml.ColorConfig{
			Color:             "blue",
			BgColor:           "black",
			TitleColor:        "red",
			TitleBgColor:      "white",
			CommentColor:      "gray",
			TitleCommentColor: "black",
		},
		Stream: toml.StreamConfig{
			LogFilterShell:    "toml-log-filter",
			ErrLogFilterShell: "toml-err-filter",
		},
	}

	// Test with terminal true and false
	for _, isTerminal := range []bool{true, false} {
		res := Convert(ctrlModel, stModels, yomelToml, isTerminal)
		if res.Ctrl.Title != "Test Title" {
			t.Errorf("expected Title to be 'Test Title', got '%s'", res.Ctrl.Title)
		}
		if len(res.Stages) != 1 {
			t.Errorf("expected 1 stage, got %d", len(res.Stages))
		}
	}
}

func TestDecideParameterOrToml(t *testing.T) {
	if got := decideParameterOrToml("", "toml"); got != "toml" {
		t.Errorf("expected 'toml', got '%s'", got)
	}
	if got := decideParameterOrToml("para", "toml"); got != "para" {
		t.Errorf("expected 'para', got '%s'", got)
	}
}

func TestToLowerWithSpaces(t *testing.T) {
	input := "TestCommentWordCase"
	expected := "test comment word case"
	if got := toLowerWithSpaces(input); got != expected {
		t.Errorf("expected '%s', got '%s'", expected, got)
	}

	// Test without upper case change
	input2 := "test"
	if got := toLowerWithSpaces(input2); got != "test" {
		t.Errorf("expected 'test', got '%s'", got)
	}
}
