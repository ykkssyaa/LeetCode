package org.LeetCode;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    private Stream<Arguments> romanParameters() {
        return Stream.of(
                Arguments.of(1, "I")

        );
    }

    @ParameterizedTest
    @MethodSource("romanParameters")
    void romanToInt(int num, String roman) {

        assertEquals(num, Solution.romanToInt(roman));

    }
}