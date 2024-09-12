package evaluate

import (
	"testing"
)

func Test_evaluator_test(t *testing.T) {
	tests := []struct {
		name string
		t    Evaluator
		args float64
		want Status
	}{
		{
			name: "TestWarningLessThanWarn",
			t:    "-w<10",
			args: 5,
			want: StatusWarning,
		},
		{
			name: "TestWarningLessThanOK",
			t:    "-w<10",
			args: 15,
			want: StatusOK,
		},
		{
			name: "TestWarningMoreThanWarn",
			t:    "-w>10",
			args: 15,
			want: StatusWarning,
		},
		{
			name: "TestWarningMoreThanOK",
			t:    "-w>10",
			args: 5,
			want: StatusOK,
		},
		{
			name: "TestWarningEqualWarn",
			t:    "-w=10",
			args: 10,
			want: StatusWarning,
		},
		{
			name: "TestWarningNotEqualWarn",
			t:    "-w!10",
			args: 5,
			want: StatusWarning,
		},
		{
			name: "TestWarningNotEqualWarn",
			t:    "-w!10",
			args: 15,
			want: StatusWarning,
		},
		{
			name: "TestWarningNotEqualOK",
			t:    "-w!=10",
			args: 10,
			want: StatusOK,
		},
		{
			name: "TestWarningEqualOK",
			t:    "-w=10",
			args: 5,
			want: StatusOK,
		},
		{
			name: "TestWarningEqualOK",
			t:    "-w==10",
			args: 5,
			want: StatusOK,
		},
		{
			name: "TestWarningLessThanEqualWarn",
			t:    "-w<=10",
			args: 10,
			want: StatusWarning,
		},
		{
			name: "TestWarningLessThanEqualWarn",
			t:    "-w<=10",
			args: 5,
			want: StatusWarning,
		},
		{
			name: "TestWarningLessThanEqualOK",
			t:    "-w<=10",
			args: 15,
			want: StatusOK,
		},
		{
			name: "TestWarningMoreThanEqualWarn",
			t:    "-w=>10",
			args: 10,
			want: StatusWarning,
		},
		{
			name: "TestWarningMoreThanEqualWarn",
			t:    "-w>=10",
			args: 15,
			want: StatusWarning,
		},
		{
			name: "TestWarningMoreThanEqualOK",
			t:    "-w>=10",
			args: 5,
			want: StatusOK,
		},
		{
			name: "TestUnknown",
			t:    "-h/10",
			args: 5,
			want: StatusUnknown,
		},
		{
			name: "TestWarningMoreThanLessThanEqualOK",
			t:    "-w>10,<20",
			args: 5,
			want: StatusOK,
		},
		{
			name: "TestWarningMoreThanLessThanEqualOK",
			t:    "-w>10,<20",
			args: 25,
			want: StatusOK,
		},
		{
			name: "TestWarningMoreThanLessThanEqualWarning",
			t:    "-w>10,<20",
			args: 15,
			want: StatusWarning,
		},
		{
			name: "TestWarningLessThanMoreThanEqualOK",
			t:    "-w<10,>20",
			args: 15,
			want: StatusOK,
		},
		{
			name: "TestWarningLessThanMoreThanEqualWarning",
			t:    "-w<10,>20",
			args: 25,
			want: StatusWarning,
		},
		{
			name: "TestWarningLessThanMoreThanEqualWarning",
			t:    "-w<10,>20",
			args: 5,
			want: StatusWarning,
		},
		{
			name: "TestWarningCriticalLessThanCrit",
			t:    "-w<20 -c<10",
			args: 5,
			want: StatusCritical,
		},
		{
			name: "TestWarningCriticalLessThanWarn",
			t:    "-w<20 -c<10",
			args: 15,
			want: StatusWarning,
		},
		{
			name: "TestWarningCriticalLessThanOK",
			t:    "-w<20 -c<10",
			args: 25,
			want: StatusOK,
		},
		{
			name: "TestWarningCriticalLessThanOK",
			t:    "-w>15,<20 -c<10",
			args: 12,
			want: StatusOK,
		},
		{
			name: "TestWarningCriticalLessThanWarn",
			t:    "-w>15,<20 -c<10",
			args: 16,
			want: StatusWarning,
		},
		{
			name: "TestWarningCriticalLessThanCrit",
			t:    "-w>15,<20 -c<10",
			args: 8,
			want: StatusCritical,
		},
		{
			name: "TestWarningMoreThanNotEqualUnknown",
			t:    "-w>15,!=20",
			args: 18,
			want: StatusUnknown,
		},
		{
			name: "TestWarningMoreThanLessThanEqualWarn",
			t:    "-w>20,<23",
			args: 22.39,
			want: StatusWarning,
		},
		{
			name: "TestCriticalWarningMoreThanCritical",
			t:    "-c<10 -w<20",
			args: 8,
			want: StatusCritical,
		},
		{
			name: "TestCriticalWarningMoreThanCritical",
			t:    "-c<10 -w<20",
			args: 15,
			want: StatusWarning,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Test(tt.args); got != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
