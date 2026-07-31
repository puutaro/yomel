package sh

// func Test_Exec(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		yomelInfos []StageInfo
// 		wantSubstr string
// 	}{
// 		{
// 			name: "should execute simple echo command and complete without error",
// 			yomelInfos: []StageInfo{
// 				{
// 					No:      1,
// 					Desc:    "echo-stage",
// 					CmdStrs: "echo 'hello yomel'",
// 					IsLog:   false,
// 				},
// 			},
// 			wantSubstr: "",
// 		},
// 		{
// 			name: "should execute multi-stage pipeline and print log when IsLog is true",
// 			yomelInfos: []StageInfo{
// 				{
// 					No:        1,
// 					Desc:      "source-stage",
// 					CmdStrs:   "echo 'data1\ndata2'",
// 					IsLog:     true,
// 					LogFilter: "grep 'data1'",
// 				},
// 			},
// 			wantSubstr: "YOMEL LOG",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Capture stdout and stderr to prevent pollution during test runs
// 			oldStdout := os.Stdout
// 			oldStderr := os.Stderr
// 			_, wOut, _ := os.Pipe()
// 			rErr, wErr, _ := os.Pipe()
// 			os.Stdout = wOut
// 			os.Stderr = wErr

// 			Exec(tt.yomelInfos)

// 			wOut.Close()
// 			wErr.Close()
// 			os.Stdout = oldStdout
// 			os.Stderr = oldStderr

// 			var bufErr bytes.Buffer
// 			_, _ = io.Copy(&bufErr, rErr)
// 			errOutput := bufErr.String()

// 			if tt.wantSubstr != "" {
// 				assert.True(t, strings.Contains(errOutput, tt.wantSubstr))
// 			}
// 		})
// 	}
// }
