package greeter

import (
    "fmt"
    "testing"
)

func TestGreeter(t *testing.T) {
    type TestResults struct {
        name string
    }
    type TestParameters struct {
        greeterName string
        // parameters
        beginFn func(t *testing.T)
        expectFn func(t *testing.T, results TestResults)
    }
    type TestScenario struct {
        name string
        setupFn func(t *testing.T) TestParameters
    }
    testScenarios := []TestScenario{
        {
            setupFn: func(t *testing.T) TestParameters {
                greeterName := "foo"
                return TestParameters{
                    greeterName: greeterName,
                    beginFn: func(t *testing.T) {},
                    expectFn: func(t *testing.T, results TestResults) {
                        expectedGreeting := fmt.Sprintf(greetingFormatter, greeterName)
                        if results.name != expectedGreeting {
                            t.Errorf(
                                "Unexpected greeting: \"%s\" != \"%s\"",
                                results.name,
                                expectedGreeting,
                            )
                        }
                    },
                }
            },
        },
    }
    for _, scenario := range testScenarios {
        parameters := scenario.setupFn(t)
        t.Run(
            scenario.name,
            func(t *testing.T) {
                parameters.beginFn(t)
                
                result := Greet(parameters.greeterName)

                parameters.expectFn(
                    t,
                    TestResults{
                        name: result,
                    },
                )
            },
        )
    }
}
